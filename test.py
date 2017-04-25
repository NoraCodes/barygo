#!/usr/bin/env python3
# A script to test BaryPy's speed. This script will generate a list of bodies of length sys.argv[1].

import sys
import random

def main():
	n = int(sys.argv[1])
	random.seed()
	for i in range(n):
		x = random.uniform(-90,110)
		y = random.uniform(-110,90)
		z = random.uniform(-100,100)
		mass = random.uniform(0.1,10)
		print(str(x) + ":" + str(y) + ":" + str(z) + ":" + str(mass))

if __name__ == "__main__":

	main()