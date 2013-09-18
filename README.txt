A neat little bit of Python+Bash hacking which renders pretty backgrounds on your iTerm2 windows, detailing the current working directory.

Problems: Needs to show active SSH connections, and needs to go much much faster.

Installation:

    pip install pillow
    git clone git://github.com/zephod/itermbg ~/.itermbg
    source ~/.itermbg/install

You can add that last line to your .bashrc file, and the effect will take place automatically, everywhere.


Addendum
========

To take this further, it might be interesting to implement a Python daemon which presents a named pipe as a FIFO queue. iTerm can be instructed to consume its background from the FIFO queue. See: http://stackoverflow.com/questions/3806210/python-interprocess-querying-control

Also this:
http://kpumuk.info/mac-os-x/how-to-show-ssh-host-name-on-the-iterms-background/
