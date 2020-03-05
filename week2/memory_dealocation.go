
// we need to deallocate the mmemory
// other wise free theemmory after we use it
// other wise we will eventually run out of memory
// becaus everytime you execute a variable
// you get a new copy of the value
// its called memory leak
// WHY WE WORRIED ABOUT THIS MEMORY LEAK SO MUCH
// we never worried abou it in python language
// YOU NEED TO UNDERSTAND A LITTLE BIT ABOUT THE

// 1) HEAP MEMORY

		// just to remember clearly anythig that is in the glocal variable is called 
		// in the heap memory
		// heap is the global and presistance reagion of memory
		// you have to explecitly deallocated in the memoty 
		// in the language like C
		// like 
		// in C you allocate memory with
		// x = malloc(32)
		// and deallocate
		// free(x)


// 2) STACK MEMORY
		// stack is a kind of memory that is dedicated to function calls
		// any thing or variable that inside function is in the stack memory
		// local variabe is in the stack
		// the stack memory is automatically deallocated
		// when the function called is finished working

//in the interpreted language it is done by the interpreter
// but it is slow
// in the compiled language it is very fast if you manually do it










//----------------------------------------