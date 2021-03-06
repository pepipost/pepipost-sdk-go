/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for ActionEnum enum
 */
type ActionEnum int

/**
 * Value collection for ActionEnum enum
 */
const (
    Action_INCREASE ActionEnum = 1 + iota
    Action_DECREASE
)

func (r ActionEnum) MarshalJSON() ([]byte, error) { 
    s := ActionEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *ActionEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  ActionEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts ActionEnum to its string representation
 */
func ActionEnumToValue(actionEnum ActionEnum) string {
    switch actionEnum {
        case Action_INCREASE:
    		return "increase"		
        case Action_DECREASE:
    		return "decrease"		
        default:
        	return "increase"
    }
}

/**
 * Converts ActionEnum Array to its string Array representation
*/
func ActionEnumArrayToValue(actionEnum []ActionEnum) []string {
    convArray := make([]string,len( actionEnum))
    for i:=0; i<len(actionEnum);i++ {
        convArray[i] = ActionEnumToValue(actionEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func ActionEnumFromValue(value string) ActionEnum {
    switch value {
        case "INCREASE":
            return Action_INCREASE
        case "DECREASE":
            return Action_DECREASE
        default:
            return Action_INCREASE
    }
}
