#include <stdlib.h>
#include <string.h>
#include "hotspot.h"

#define SEPARATORS " \t"

size_t parseParams(str_pair *table, size_t max, const char *params) {
	char name = 1;
	char *save;
	size_t count = 0;

	char *p = strtok_r((char *)params, SEPARATORS, &save);
	while (p != NULL && count < max) {
		if (name) strcpy(table[count].name, p);
		else strcpy(table[count++].value, p);
		p = strtok_r(NULL, SEPARATORS, &save);
		name = !name;
	}

	return count;
}

HotSpot *newHotSpot(const char *floorplan, const char *config, const char *params) {
	HotSpot *h = (HotSpot *)malloc(sizeof(HotSpot));
	h->config = default_thermal_config();

	str_pair *table = (str_pair *)malloc(sizeof(str_pair) * MAX_ENTRIES);

	if (config && strlen(config) > 0) {
		size_t count = read_str_pairs(table, MAX_ENTRIES, (char *)config);
		thermal_config_add_from_strs(&h->config, table, count);
	}

	if (params && strlen(params) > 0) {
		size_t count = parseParams(table, MAX_ENTRIES, params);
		thermal_config_add_from_strs(&h->config, table, count);
	}

	free(table);

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

void copyC(const HotSpot *h, double *dst) {
	const double *src = h->model->block->a;
	size_t i, count = h->nodes;

	for (i = 0; i < count; i++)
		dst[i] = src[i];
}

void copyG(const HotSpot *h, double *dst) {
	const double *const *src = (const double *const *)h->model->block->b;
	size_t i, j, count = h->nodes;

	for (i = 0; i < count; i++)
		for (j = 0; j < count; j++)
			dst[i * count + j] = src[i][j];
}
