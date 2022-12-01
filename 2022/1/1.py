with open('input.txt') as input:
	current = 0
	topthree = list()
	for line in input.readlines():
		line = line.strip()
		if line:
			if line.isnumeric():
				current += int(line)
		else:
			topthree.append(current)
			topthree.sort(reverse=True)
			topthree = topthree[0:3]
			current = 0
print("final max: " + str((topthree[0] + topthree[1] + topthree[2])))