
def l337(input_string):
    replacement_table = {
      "o": "0",
      "i": "1",
      "z": "2",
      "e": "3",
      "a": "4",
      "s": "5",
      "g": "6",
      "t": "7",
      }
    output_string = ""
    for char in input_string:
        if char in replacement_table:
            output_string += replacement_table[char]
        else:
            output_string += char
    return output_string


print(l337("leet"))
print(l337("noob"))
print(l337("skiddie"))
print(l337("zip"))
print(l337("gateway"))

