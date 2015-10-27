#include <stdlib.h>
#include "circuit.h"

Circuit *newCircuit(const char *floorplan, const char *config) {
	thermal_config_t thermal_config = default_thermal_config();

	if (config && strlen(config) > 0) {
		str_pair *table = (str_pair *)malloc(sizeof(str_pair) * MAX_ENTRIES);
		if (!table) goto err_malloc_str_pair;

		size_t count = read_str_pairs(table, MAX_ENTRIES, (char *)config);
		thermal_config_add_from_strs(&thermal_config, table, count);

		free(table);
	}

	flp_t *flp = read_flp((char *)floorplan, FALSE);
	if (!flp) goto err_read_flp;

	RC_model_t *model = alloc_RC_model(&thermal_config, flp, 0);
	if (!model) goto err_alloc_RC_model;
	if (model->type != BLOCK_MODEL) goto err_model_type;

	populate_R_model(model, flp);
	populate_C_model(model, flp);

	Circuit *circuit = (Circuit *)malloc(sizeof(Circuit));
	if (!circuit) goto err_malloc_Circuit;

	circuit->units = flp->n_units;
	circuit->nodes = model->block->n_nodes;

	size_t i, j, nodes = circuit->nodes;

	circuit->capacitance = (double *)malloc(nodes * sizeof(double));
	if (!circuit->capacitance) goto err_malloc_capacitance;

	for (i = 0; i < nodes; i++)
		circuit->capacitance[i] = model->block->a[i];

	circuit->conductance = (double *)malloc(nodes * nodes * sizeof(double));
	if (!circuit->conductance) goto err_malloc_conductance;

	for (i = 0; i < nodes; i++)
		for (j = 0; j < nodes; j++)
			circuit->conductance[i * nodes + j] = model->block->b[i][j];

	delete_RC_model(model);
	free_flp(flp, FALSE);

	return circuit;

err_malloc_conductance:
	free(circuit->capacitance);

err_malloc_capacitance:
	free(circuit);

err_malloc_Circuit:
err_model_type:
	delete_RC_model(model);

err_alloc_RC_model:
	free_flp(flp, FALSE);

err_read_flp:
err_malloc_str_pair:

	return NULL;
}

void dropCircuit(Circuit *circuit) {
	free(circuit->capacitance);
	free(circuit->conductance);
	free(circuit);
}
