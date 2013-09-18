#!/usr/bin/env python

from PIL import Image, ImageDraw, ImageFont
import subprocess
import os
import sys

CHAR_WIDTH = 14
CHAR_HEIGHT = 34

RAW_APPLESCRIPT = """
    tell application "iTerm"
      repeat with theTerminal in terminals
        tell theTerminal
          try
            tell session id "%s"
              set background image path to "%s"
            end tell
          on error errmesg number errn
          end try
        end tell
      end repeat
    end tell
"""

def my_dir():
    path = os.path.dirname(os.path.abspath(__file__))
    return path

def gen_img(line1,line2=None):
    _filename = hash(line1+'::'+str(line2))
    filename = os.path.join(my_dir(),'cache',str(_filename))
    if os.path.exists(filename):
        return filename
    print 'generating'

    cols,lines = get_size()
    w = cols * CHAR_WIDTH * 2
    h = lines * CHAR_HEIGHT * 2

    img = Image.new('RGBA',(w, h))
    draw = ImageDraw.Draw(img)
    size = 50 * 2

    f = ImageFont.truetype(os.path.join(my_dir(),'Ubuntu-R.ttf'),95)
    txt_w, txt_h = draw.textsize(line1, f)
    draw.text( (w-txt_w-40, 40), line1, fill=(255,100,100), font=f)

    img.save(filename, 'GIF', transparency=1)
    return filename

def set_img(tty,img):
    applescript = RAW_APPLESCRIPT % (tty,img)
    process = subprocess.Popen(['osascript', '-e', applescript], stdout=subprocess.PIPE)
    stdout,stderr = process.communicate()

def get_tty():
    p = subprocess.Popen(['tty'],stdout=subprocess.PIPE)
    stdin,stderr = p.communicate()
    return stdin[:-1]

def get_size():
    p1 = subprocess.Popen(['tput','cols'],stdout=subprocess.PIPE)
    p1_in,p1_err = p1.communicate()
    p2 = subprocess.Popen(['tput','lines'],stdout=subprocess.PIPE)
    p2_in,p2_err = p2.communicate()
    cols = int(p1_in[:-1])
    lines = int(p2_in[:-1])
    return cols,lines

if __name__=='__main__':
    home = os.path.expanduser('~')
    line1 = os.getcwd().replace(home,'~')
    print 'overwriting bg'
    filename = gen_img( line1 )
    set_img(get_tty(),filename)
