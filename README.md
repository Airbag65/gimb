# Gimb
```Gimb``` is a very scaled down clone of ```Vim```, written in ```Golang``` as a small project to learn and get comfortable with the language.
The name is a play on the three words that make up the projekt: ```Go```, ```Vim``` and ```dumb```.  

Gimb is still very much in development and can not really be used for anything yet.

### How to use Gimb
If you for some unexplainable reason would like to try and use Gimb, this following section tells you how to set it up  

As of now, Gimb comes only with a Makefile that cleans up the go module, builds the project and moves the binary to the home folder, i.e. ```~/```. 
To use it system wide you'll need to make an alias in your ```.$(SHELL)rc``` file. Mine is as follows:
```Bash
    alias gimb="~/gimb_out"
```

Once inside of ```GIMB``` you can use the most basic Vim-bindings to move around. ```h``` to move left, ```j``` to move down, ```k``` to move up and ```l``` to move right.
You can use ```:``` to enter COMMAND-Mode. Here you can use the most essential commands also available in VIM, such as ```q```, ```w```, ```wq```. You can also use COMMAND
```h``` or ```help``` to get an overview of the available commands. To enter INSERT-Mode, type ```i``` when in NORMAL-Mode. To exit INSERT-Mode, press ```ESC```
