# Go Pokedex CLI

Go Pokedex CLI is a REPL command line utility written entirely in Go that communicates with the [PokéAPI](https://pokeapi.co/) to provide a fun, interactive Pokémon experience.

## Command History Provided by the `Liner` package
The Go Pokedex CLI tool leverages the third-party [Liner](https://github.com/peterh/liner) package to enable clean REPL input with command history accessible via the Up and Down arrow keys.

## Supported Commands and Examples

### HELP
The `help` command displays basic information on the available commands.

Example:
```
Pokedex > help

Welcome to the Pokedex!
Usage:

exit: Exits the Pokedex
map: Displays the next 20 names of location areas in the Pokemon world
mapb: Displays the previous 20 names of location areas in the Pokemon world
explore: Shows a list of Pokemon available in a given location
catch: Attempts to catch a pokemon by name
inspect: Gets the details of a Pokemon that you've already caught
pokedex: Prints a list of all Pokemon you have caught
help: Displays a help message
```

### EXIT
The `exit` command exits the Pokedex tool.

### MAP
The `map` tool displays 20 names of location areas in the Pokemon world. The `map` command is progressive. That is, each time you enter the `map` command, the next 20 locations are displayed. You can keep using the `map` command until you reach the end of the list of locations supplied by PokeAPI. At the time this was written, the PokeAPI held **over 1000** locations!

Example:
```
Pokedex > map

canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

### MAPB
The `mapb` command functions similarly to the `map` command, except it advances **backwards** through the list of locations. This command will not work until you run `map` at least twice.

Examples:
```
Pokedex > mapb

no previous location available - try using the `map` command first
```

```
Pokedex > mapb

mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
great-marsh-area-3
great-marsh-area-4
great-marsh-area-5
great-marsh-area-6
solaceon-ruins-2f
solaceon-ruins-1f
solaceon-ruins-b1f-a
solaceon-ruins-b1f-b
solaceon-ruins-b1f-c
solaceon-ruins-b2f-a
solaceon-ruins-b2f-b
solaceon-ruins-b2f-c
solaceon-ruins-b3f-a
solaceon-ruins-b3f-b
solaceon-ruins-b3f-c
```

### EXPLORE
The `explore` command takes a single location name as an argument. It returns the Pokemon available to be discovered in that area. You can use any of the locations returned by the `map` command.

Example:
```
Pokedex > explore canalave-city-area

Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados
 - wingull
 - pelipper
 - shellos
 - gastrodon
 - finneon
 - lumineon
```

### CATCH
The `catch` command takes a single pokemon name as an argument. It then attempts to *catch* that pokemon using a random number generator. If the catch attempt fails, the program will tell you that the pokemon escaped. If the catch attempt succeeds, the pokemon's information will be saved to your pokedex, which can be viewed using the `pokedex` command.

Example:
```
Pokedex > catch squirtle

Throwing a Pokeball at squirtle...
squirtle escaped!

Pokedex > catch squirtle

Throwing a Pokeball at squirtle...
squirtle was caught!
Adding squirtle to your pokedex...
```

### INSPECT
The `inspect` command takes a single pokemon name as an argument. It can be used on any pokemon that you have already caught. It prints out some basic details on the pokemon.

Example:
```
Pokedex > inspect squirtle

Name: squirtle
Height: 5
Weight: 90
Stats:
 -hp: 44
 -attack: 48
 -defense: 65
 -special-attack: 50
 -special-defense: 64
 -speed: 43
Types:
 - water
```

### POKEDEX
The `pokedex` command accepts no arguments. It prints out a list of all pokemon already caught. If no pokemon are found in the pokedex, an error will print instead.

Example:
```
Pokedex > pokedex

Your Pokedex:
 - pikachu
 - squirtle
```