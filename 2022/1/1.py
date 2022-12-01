with open('input.txt') as input:
	cur = 0
	max = 0
	for line in input.readlines():
		line = line.strip()
		if line:
			if line.isnumeric():
				cur += int(line)
		else:
			if max < cur:
				max = cur
			cur = 0
	if max < cur:
		max = cur

print("final max: " + str(max))