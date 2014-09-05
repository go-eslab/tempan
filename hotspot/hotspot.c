#include <stdlib.h>
#include <string.h>
#include "hotspot.h"

#define WHITESPACE " \t"

size_t parseConfigLine(str_pair *table, size_t max, const char *line) {
	char first = 1;
	size_t count = 0;

	char *p = strtok((char *)line, WHITESPACE);
	while (p != NULL && count < max) {
		if (first) strcpy(table[count].value, p);
		else strcpy(table[count++].name, p);
		first = !first;
	}

	return count;
}

HotSpot *newHotSpot(const char *floorplan, const char *config, const char *line) {
	HotSpot *h = (HotSpot *)malloc(sizeof(HotSpot));

	h->config = default_thermal_config();

	if (config && strlen(config) > 0) {
		str_pair table[MAX_ENTRIES];
		size_t count = read_str_pairs(table, MAX_ENTRIES, (char *)config);
		thermal_config_add_from_strs(&h->config, table, count);
	}

	if (line && strlen(line) > 0) {
		str_pair table[MAX_ENTRIES];
		size_t count = parseConfigLine(table, MAX_ENTRIES, line);
		thermal_config_add_from_strs(&h->config, table, count);
	}

	h->floorplan = read_flp((char *)floorplan, FALSE);
	h->model = alloc_RC_model(&h->config, h->floorplan);

	populate_R_model(h->model, h->floorplan);
	populate_C_model(h->model, h->floorplan);

	h->cores = h->floorplan->n_units;
	h->nodes = h->model->block->n_nodes;

	return h;
}

void freeHotSpot(HotSpot *h) {
	delete_RC_model(h->model);
	free_flp(h->floorplan, FALSE);
	free(h);
}

void copyA(const HotSpot *h, double *dst) {
	const double *src = h->model->block->a;
	size_t count = h->nodes;

	for (size_t i = 0; i < count; i++)
		dst[i] = src[i];
}

void copyB(const HotSpot *h, double *dst) {
	const double *const *src = (const double *const *)h->model->block->b;
	size_t count = h->nodes;

	for (size_t i = 0; i < count; i++)
		for (size_t j = 0; j < count; j++)
			dst[i * count + j] = src[i][j];
}

void copyG(const HotSpot *h, double *dst) {
	const double *const *src = (const double *const *)h->model->block->g;
	size_t count = h->nodes;

	for (size_t i = 0; i < count; i++)
		for (size_t j = 0; j < count; j++)
			dst[i * count + j] = src[i][j];
}
