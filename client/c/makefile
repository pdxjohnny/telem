CC=gcc
CFLAGS=-c -Wall
LDFLAGS=$(shell gpgme-config --libs)
LDFLAGS+=$(shell gpg-error-config --libs)
LDFLAGS+=$(shell libassuan-config --libs)
LDFLAGS+=--static
SOURCES=$(wildcard *.c)
OBJECTS=$(SOURCES:.c=.o)
EXECUTABLE=test-client

all: $(SOURCES) $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
	$(CC) $(OBJECTS) $(LDFLAGS) -o $@

.c.o:
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm *.o $(EXECUTABLE)
