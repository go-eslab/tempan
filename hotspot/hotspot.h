#ifndef __HOTSPOT_H__
#define __HOTSPOT_H__

#include <util.h>
#include <flp.h>
#include <temperature.h>
#include <temperature_block.h>

typedef struct {
	thermal_config_t config;
	flp_t            *floorplan;
	RC_model_t       *model;
	size_t           nodes;
	size_t           cores;
} HotSpot;

HotSpot *newHotSpot(const char *floorplan, const char *config, const char *params);
void freeHotSpot(HotSpot *);

void copyA(const HotSpot *, double *);
void copyB(const HotSpot *, double *);
void copyG(const HotSpot *, double *);

#endif
