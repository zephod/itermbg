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

def gen_img(line1,line2=None,pattern=None):
    _filename = hash(line1+'::'+str(line2))
    filename = os.path.join(my_dir(),'cache',str(_filename))

    #if os.path.exists(filename):
    #    return filename

    cols,lines = get_size()
    w = cols * CHAR_WIDTH + 25
    h = lines * CHAR_HEIGHT + 8

    img = Image.new('RGBA',(w, h))
    draw = ImageDraw.Draw(img)
    size = 50 * 2
    
    # Blit the background pattern onto the image
    if pattern:
        pattern = Image.open(pattern)
        pw,ph = pattern.size
        for y in range(0,h+ph,ph):
            for x in range(0,w+pw,pw):
                img.paste( pattern, (x,y))

    f = ImageFont.truetype(os.path.join(my_dir(),'/Library/Fonts/Microsoft/Matura Script Capitals'),65)
    txt_w, txt_h = draw.textsize(line1, f)
    draw.text( (w-txt_w-10, 20), line1, fill=(100,100,100), font=f)

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
    pattern = '/Users/zephod/.itermbg/carbon.png'
    pattern = '/Users/zephod/.itermbg/honey.png'
    pattern = '/Users/zephod/.itermbg/microcarbon.png'
    #pattern = '/Users/zephod/.itermbg/bindingdark.png'
    #pattern = '/Users/zephod/.itermbg/maze.png'
    #pattern = '/Users/zephod/.itermbg/tweed.png'
    filename = gen_img( line1,pattern=pattern )
    set_img(get_tty(),filename)
