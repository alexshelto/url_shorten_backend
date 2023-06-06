package BaseConversion


/*
465
465 / 62 = 7 R 31  ("7")
31                 ("v")
*/
func ConvertToBase62(number int) string {
    const BASE = 62
    ret := ""
    for {
        if number >= BASE {
            based := number / BASE
            number = number % BASE
            ret += Base62Mappings[based]
        } else {
            ret += Base62Mappings[number]
            return ret
        } 
    }
}

/*

0: a,
1: b,
2: c, 
3: d,
4: e,
5: f, 
6: g,
7: h, 
8: i,
9: j, 
10: k
*/

/*
0:

// what i want is...
22 22 / 10: 2 (c) 
2             (c)
22 = cc

41: 
41 / 10 = 4 R 1 (4) => e
1               (1) => b


rules...  
if num > base..  
based = math.floor(num/base) 
num - base 
map[based] 


63: 
63 / 62 = 1 R 1 ("1")
1               ("1")

64
64 / 62 = 1 R 2 ("1")
2               ("2")

465
465 / 62 = 7 R 31  ("7")
31                 ("v")

{1, "1"},
{4, "4"},
{39,"C"},
{62, "0"},
{63, "11"},
{64, "12"},
{465, "7v"},
*/

var Base62Mappings = map[int]string{
    0: "0",
    1: "1",
    2: "2",
    3: "3", 
    4: "4", 
    5: "5", 
    6: "6", 
    7: "7", 
    8: "8", 
    9: "9",
    10: "a",
    11: "b",
    12: "c",
    13: "d",
    14: "e",
    15: "f",
    16: "g",
    17: "h",
    18: "i",
    19: "j",
    20: "k",
    21: "l",
    22: "m",
    23: "n",
    24: "o",
    25: "p",
    26: "q",
    27: "r",
    28: "s",
    29: "t",
    30: "u",
    31: "v",
    32: "w",
    33: "x",
    34: "y",
    35: "z",
    36: "A",
    37: "B",
    38: "C",
    39: "D",
    40: "E",
    41: "F",
    42: "G",
    43: "H",
    44: "I",
    45: "J",
    46: "K",
    47: "L",
    48: "M",
    49: "N",
    50: "O",
    51: "P",
    52: "Q",
    53: "R",
    54: "S",
    55: "T",
    56: "U",
    57: "V",
    58: "W",
    59: "X",
    60: "Y",
    61: "Z",

}

