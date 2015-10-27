#ifndef __CIRCUIT_H__
#define __CIRCUIT_H__

#include <util.h>
#include <flp.h>
#include <temperature.h>
#include <temperature_block.h>

typedef struct {
	size_t units;
	size_t nodes;
	double *capacitance;
	double *conductance;
} Circuit;

Circuit *newCircuit(const char *, const char *);
void dropCircuit(Circuit *);

#endif
