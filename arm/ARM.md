## ARM Support

I've been running this Discord bot on a Raspberry PI3-b model - which is an ARM 32 bit processor. Not x86_64 like my personal computer. Therefore I must cross-compile if I am planning on compiling on my personal computer and then running the binary on the raspberry.

### Script

Run the cross-compile script `pi_compile.sh` to compile the program in `arm32`.

- `./pi_compile.sh`

### Moving to Pi 

You can copy the program to the Pi when on the same network using `scp`.

`scp spin_bot_arm64 <username>@<PI IP or hostname.local:/<path-on-pi>`

N.B to use `hostname.local` (where `hostname` is the hostname of the Pi) you must have multi-cast DNS setup on host machine e.g. via Avahi. Otherwise you need to know the IP address of the Pi on your local network.

### Running on Pi 

Once moved, you can run the program on Pi just as you would on your normal computer. However, you may want to consider using `nohup` to prevent the bot from stopping once you disconnect your session and sending the stdout to a logfile.

`nohup ./spin_bot_arm64 > logfile &`

