#include <stdio.h>
#include <string.h>

int main()
{
	FILE * input;
	input = fopen("input","r");
	
	int trees = 0;
	int pointer = 0;
	
	char line[64];
	char current;
	
	while(fgets(line,sizeof(line),input))
	{
		current = line[pointer];
		printf("%s",line);
		
		for(int i = pointer;i>0;i--) printf(" ");
		printf("^%i\n",pointer);
		
		if(current == '#') trees++;
		
		pointer += 3;
		if(pointer > 30) pointer -= 31;
	}
	
	printf("%i\n",trees);
	return trees;
}
