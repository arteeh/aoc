#include <stdio.h>
#include <string.h>

int t(int right, int down)
{
	int trees = 0;
	int pointer = 0;
	char line[33];
	
	FILE * input;
	input = fopen("input","r");
	while(fgets(line,sizeof(line),input))
	{
		//printf("%s",line);
		//for(int i = pointer;i>0;i--) printf(" ");
		//printf("^%i\n",pointer);
		
		if(line[pointer] == '#') trees++;
		
		pointer += right;
		if(pointer > 30) pointer -= 31;
		for(int i = 1; i < down; i++) fgets(line,sizeof(line),input);
	}
	fclose(input);
	
	printf("%i,%i:\t%i\n",right,down,trees);
	return trees;
}

int main()
{
	long total = 1;
	
	total *= t(1,1);
	total *= t(3,1);
	total *= t(5,1);
	total *= t(7,1);
	total *= t(1,2);

	printf("total:\t%li\n",total);
	return total;
}
