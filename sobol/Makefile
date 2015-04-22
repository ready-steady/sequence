objects := sobol.o

all: $(objects)

%.o: %.cpp %.hpp
	g++ -c -g $<

clean:
	$(RM) -f $(objects)
