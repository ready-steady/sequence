#include <string>

using namespace std;

#include "sobol/sobol.hpp"

extern "C" {
	void sobol(int dim_num, long long int *seed, double quasi[]) {
		i8_sobol(dim_num, seed, quasi);
	}
}
