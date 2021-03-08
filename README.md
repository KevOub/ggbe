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