// generated by stringer -type=VirtualCategory; DO NOT EDIT

package events

import "fmt"

const _VirtualCategory_name = "TYPINGPOINTINGWINDOWING"

var _VirtualCategory_index = [...]uint8{0, 6, 14, 23}

func (i VirtualCategory) String() string {
	if i+1 >= VirtualCategory(len(_VirtualCategory_index)) {
		return fmt.Sprintf("VirtualCategory(%d)", i)
	}
	return _VirtualCategory_name[_VirtualCategory_index[i]:_VirtualCategory_index[i+1]]
}
