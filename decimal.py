with open("subaru-sas-swapped-order.bin", "rb") as input_file:
    bytes_data = input_file.read()

word_data_swapped = []

for i in range(int(len(bytes_data)/2)):
    bytes_pair = bytes_data[i*2:i*2+2]
    print(int.from_bytes(bytes_pair, byteorder='big', signed=False))
