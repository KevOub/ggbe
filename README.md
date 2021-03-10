# ggbe
Golang Gameboy Emulator

## Gameboy opcode table
Got from https://github.com/lmmendes/game-boy-opcodes/

How the debugger works: 
JSON -> Hashmap
This hashmap's keys correspond to mnemonic, length, cycles, flags, address, etc.
Then, added a DEBUG flag to the gcpu.go 



### TODO
0) Cartridge loader
1) Implement JSON lookup for opcode table to show the corresponding command with missing hex 


--- 
2) Implement missing opcodes. Include loggin 
---
3) Implement timing
---
4) Implement User Input
---
5) Implement Graphics
---
6)

Resources:

https://gekkio.fi/files/gb-docs/gbctr.pdf

## generating large switch case
bash command "seq 0 255 | while read n; do printf "%02X\n" $n; done > switchcase"
vim switchcase

    Press Esc to enter 'command mode'
    Use Ctrl+V to enter visual block mode
    Move Up/Downto select the columns of text in the lines you want to comment.
    Then hit Shift+i and type the text you want to insert.
    Then hit Esc, wait 1 second and the inserted text will appear on every line.

sqc: [https://stackoverflow.com/questions/253380/how-to-insert-text-at-beginning-of-a-multi-line-selection-in-vi-vim]