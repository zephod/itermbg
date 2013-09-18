A neat little bit of Python+Bash hacking which renders pretty backgrounds on your iTerm2 windows, detailing the current working directory.

Problems: Needs to show active SSH connections, and needs to go much much faster.

Installation:

    pip install pillow
    git clone git://github.com/zephod/itermbg ~/.itermbg
    source ~/.itermbg/install

You can add that last line to your .bashrc file, and the effect will take place automatically, everywhere.
