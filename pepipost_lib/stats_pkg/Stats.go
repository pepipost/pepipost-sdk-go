/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package stats_pkg

import "time"
import "pepipost_lib/configuration_pkg"
import "pepipost_lib/models_pkg"

/*
 * Interface for the STATS_IMPL
 */
type STATS interface {
    GetStatsGET (*time.Time, *time.Time, models_pkg.AggregatedByEnum, *int64, *int64) (interface{}, error)
}

/*
 * Factory for the STATS interaface returning STATS_IMPL
 */
func NewSTATS(config configuration_pkg.CONFIGURATION) *STATS_IMPL {
    client := new(STATS_IMPL)
    client.config = config
    return client
}