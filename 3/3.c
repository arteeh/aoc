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
		printf("^\n");
		
		if(current == '#') trees++;
		pointer += 3;
		if(pointer > strlen(line)) break;
	}
	
	printf("%i\n",trees);
	return trees;
}
