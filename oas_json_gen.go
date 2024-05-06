// Code generated by ogen, DO NOT EDIT.

package policeuk

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode encodes DatasourceAvailability as json.
func (s DatasourceAvailability) Encode(e *jx.Encoder) {
	unwrapped := []DatasourceAvailabilityItem(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes DatasourceAvailability from json.
func (s *DatasourceAvailability) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DatasourceAvailability to nil")
	}
	var unwrapped []DatasourceAvailabilityItem
	if err := func() error {
		unwrapped = make([]DatasourceAvailabilityItem, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem DatasourceAvailabilityItem
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = DatasourceAvailability(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s DatasourceAvailability) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DatasourceAvailability) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *DatasourceAvailabilityItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DatasourceAvailabilityItem) encodeFields(e *jx.Encoder) {
	{
		if s.Date.Set {
			e.FieldStart("date")
			s.Date.Encode(e)
		}
	}
	{
		if s.StopAndSearch != nil {
			e.FieldStart("stop-and-search")
			e.ArrStart()
			for _, elem := range s.StopAndSearch {
				e.Str(elem)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfDatasourceAvailabilityItem = [2]string{
	0: "date",
	1: "stop-and-search",
}

// Decode decodes DatasourceAvailabilityItem from json.
func (s *DatasourceAvailabilityItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DatasourceAvailabilityItem to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "date":
			if err := func() error {
				s.Date.Reset()
				if err := s.Date.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"date\"")
			}
		case "stop-and-search":
			if err := func() error {
				s.StopAndSearch = make([]string, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem string
					v, err := d.Str()
					elem = string(v)
					if err != nil {
						return err
					}
					s.StopAndSearch = append(s.StopAndSearch, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"stop-and-search\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DatasourceAvailabilityItem")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DatasourceAvailabilityItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DatasourceAvailabilityItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}