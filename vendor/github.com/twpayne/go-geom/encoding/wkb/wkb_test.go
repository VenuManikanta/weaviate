package wkb

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkbcommon"
	"github.com/twpayne/go-geom/internal/testdata"
)

func test(t *testing.T, g geom.T, xdr []byte, ndr []byte) {
	if xdr != nil {
		if got, err := Unmarshal(xdr); err != nil || !reflect.DeepEqual(got, g) {
			t.Errorf("Unmarshal(%s) == %#v, %#v, want %#v, nil", hex.EncodeToString(xdr), got, err, g)
		}
		if got, err := Marshal(g, XDR); err != nil || !reflect.DeepEqual(got, xdr) {
			t.Errorf("Marshal(%#v, XDR) == %s, %#v, want %s, nil", g, hex.EncodeToString(got), err, hex.EncodeToString(xdr))
		}
	}
	if ndr != nil {
		if got, err := Unmarshal(ndr); err != nil || !reflect.DeepEqual(got, g) {
			t.Errorf("Unmarshal(%s) == %#v, %#v, want %#v, nil", hex.EncodeToString(ndr), got, err, g)
		}
		if got, err := Marshal(g, NDR); err != nil || !reflect.DeepEqual(got, ndr) {
			t.Errorf("Marshal(%#v, NDR) == %s, %#v, want %#v, nil", g, hex.EncodeToString(got), err, hex.EncodeToString(ndr))
		}
	}
	switch g := g.(type) {
	case *geom.Point:
		var p Point
		if xdr != nil {
			if err := p.Scan(xdr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", p, string(xdr), err)
			}
			if !reflect.DeepEqual(p, Point{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(xdr), p, Point{g})
			}
		}
		if ndr != nil {
			if err := p.Scan(ndr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", p, string(ndr), err)
			}
			if !reflect.DeepEqual(p, Point{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(ndr), p, Point{g})
			}
		}
	case *geom.LineString:
		var ls LineString
		if xdr != nil {
			if err := ls.Scan(xdr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", ls, string(xdr), err)
			}
			if !reflect.DeepEqual(ls, LineString{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(xdr), ls, LineString{g})
			}
		}
		if ndr != nil {
			if err := ls.Scan(ndr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", ls, string(ndr), err)
			}
			if !reflect.DeepEqual(ls, LineString{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(ndr), ls, LineString{g})
			}
		}
	case *geom.Polygon:
		var p Polygon
		if xdr != nil {
			if err := p.Scan(xdr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", p, string(xdr), err)
			}
			if !reflect.DeepEqual(p, Polygon{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(xdr), p, Polygon{g})
			}
		}
		if ndr != nil {
			if err := p.Scan(ndr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", p, string(ndr), err)
			}
			if !reflect.DeepEqual(p, Polygon{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(ndr), p, Polygon{g})
			}
		}
	case *geom.MultiPoint:
		var mp MultiPoint
		if xdr != nil {
			if err := mp.Scan(xdr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", mp, string(xdr), err)
			}
			if !reflect.DeepEqual(mp, MultiPoint{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(xdr), mp, MultiPoint{g})
			}
		}
		if ndr != nil {
			if err := mp.Scan(ndr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", mp, string(ndr), err)
			}
			if !reflect.DeepEqual(mp, MultiPoint{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(ndr), mp, MultiPoint{g})
			}
		}
	case *geom.MultiLineString:
		var mls MultiLineString
		if xdr != nil {
			if err := mls.Scan(xdr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", mls, string(xdr), err)
			}
			if !reflect.DeepEqual(mls, MultiLineString{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(xdr), mls, MultiLineString{g})
			}
		}
		if ndr != nil {
			if err := mls.Scan(ndr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", mls, string(ndr), err)
			}
			if !reflect.DeepEqual(mls, MultiLineString{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(ndr), mls, MultiLineString{g})
			}
		}
	case *geom.MultiPolygon:
		var mp MultiPolygon
		if xdr != nil {
			if err := mp.Scan(xdr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", mp, string(xdr), err)
			}
			if !reflect.DeepEqual(mp, MultiPolygon{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(xdr), mp, MultiPolygon{g})
			}
		}
		if ndr != nil {
			if err := mp.Scan(ndr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", mp, string(ndr), err)
			}
			if !reflect.DeepEqual(mp, MultiPolygon{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(ndr), mp, MultiPolygon{g})
			}
		}
	case *geom.GeometryCollection:
		var gc GeometryCollection
		if xdr != nil {
			if err := gc.Scan(xdr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", gc, string(xdr), err)
			}
			if !reflect.DeepEqual(gc, GeometryCollection{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(xdr), gc, GeometryCollection{g})
			}
		}
		if ndr != nil {
			if err := gc.Scan(ndr); err != nil {
				t.Errorf("%#v.Scan(%#v) == %v, want nil", gc, string(ndr), err)
			}
			if !reflect.DeepEqual(gc, GeometryCollection{g}) {
				t.Errorf("Scan(%#v) got %#v, want %#v", string(ndr), gc, GeometryCollection{g})
			}
		}
	}
}

func Test(t *testing.T) {
	for _, tc := range []struct {
		g   geom.T
		xdr []byte
		ndr []byte
	}{
		{
			g:   geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{1, 2}),
			xdr: []byte("\x00\x00\x00\x00\x01?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@"),
		},
		{
			g:   geom.NewPoint(geom.XYZ).MustSetCoords(geom.Coord{1, 2, 3}),
			xdr: []byte("\x00\x00\x00\x03\xe9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			g:   geom.NewPoint(geom.XYM).MustSetCoords(geom.Coord{1, 2, 3}),
			xdr: []byte("\x00\x00\x00\x07\xd1?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			g:   geom.NewPoint(geom.XYZM).MustSetCoords(geom.Coord{1, 2, 3, 4}),
			xdr: []byte("\x00\x00\x00\x0b\xb9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			g:   geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{{1, 2}, {3, 4}}),
			xdr: []byte("\x00\x00\x00\x00\x02\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\x02\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			g:   geom.NewLineString(geom.XYZ).MustSetCoords([]geom.Coord{{1, 2, 3}, {4, 5, 6}}),
			xdr: []byte("\x00\x00\x00\x03\xea\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xea\x03\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			g:   geom.NewLineString(geom.XYM).MustSetCoords([]geom.Coord{{1, 2, 3}, {4, 5, 6}}),
			xdr: []byte("\x00\x00\x00\x07\xd2\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xd2\x07\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			g:   geom.NewLineString(geom.XYZM).MustSetCoords([]geom.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			xdr: []byte("\x00\x00\x00\x0b\xba\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xba\x0b\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @"),
		},
		{
			g:   geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}),
			xdr: []byte("\x00\x00\x00\x00\x03\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\x03\x00\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@"),
		},
		{
			g:   geom.NewPolygon(geom.XYZ).MustSetCoords([][]geom.Coord{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}),
			xdr: []byte("\x00\x00\x00\x03\xeb\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xeb\x03\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			g:   geom.NewPolygon(geom.XYM).MustSetCoords([][]geom.Coord{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}),
			xdr: []byte("\x00\x00\x00\x07\xd3\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xd3\x07\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			g:   geom.NewPolygon(geom.XYZM).MustSetCoords([][]geom.Coord{{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}}}),
			xdr: []byte("\x00\x00\x00\x0b\xbb\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00@$\x00\x00\x00\x00\x00\x00@&\x00\x00\x00\x00\x00\x00@(\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xbb\x0b\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00$@\x00\x00\x00\x00\x00\x00&@\x00\x00\x00\x00\x00\x00(@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			g:   geom.NewMultiPoint(geom.XY).MustSetCoords([]geom.Coord{{1, 2}, {3, 4}}),
			xdr: []byte("\x00\x00\x00\x00\x04\x00\x00\x00\x02\x00\x00\x00\x00\x01?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\x04\x00\x00\x00\x02\x00\x00\x00\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			g:   geom.NewMultiPoint(geom.XYZ).MustSetCoords([]geom.Coord{{1, 2, 3}, {4, 5, 6}}),
			xdr: []byte("\x00\x00\x00\x03\xec\x00\x00\x00\x02\x00\x00\x00\x03\xe9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\xe9@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xec\x03\x00\x00\x02\x00\x00\x00\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			g:   geom.NewMultiPoint(geom.XYM).MustSetCoords([]geom.Coord{{1, 2, 3}, {4, 5, 6}}),
			xdr: []byte("\x00\x00\x00\x07\xd4\x00\x00\x00\x02\x00\x00\x00\x07\xd1?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x07\xd1@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xd4\x07\x00\x00\x02\x00\x00\x00\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			g:   geom.NewMultiPoint(geom.XYZM).MustSetCoords([]geom.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			xdr: []byte("\x00\x00\x00\x0b\xbc\x00\x00\x00\x02\x00\x00\x00\x0b\xb9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0b\xb9@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xbc\x0b\x00\x00\x02\x00\x00\x00\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @"),
		},
		{
			g: geom.NewGeometryCollection().MustPush(
				geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{-79.3698576, 43.6456613}),
				geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{{-79.3707986, 43.6453697}, {-79.3704747, 43.6454819}, {-79.370186, 43.6455592}, {-79.3699323, 43.6456385}, {-79.3698576, 43.6456613}}),
				geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{{-79.3698576, 43.6456613}, {-79.3698057, 43.6455265}}),
			),
			xdr: []byte("\x00\x00\x00\x00\x07\x00\x00\x00\x03\x00\x00\x00\x00\x01\xc0\x53\xd7\xab\xbf\x36\x0b\x55\x40\x45\xd2\xa5\x07\x8b\xe5\x7c\x00\x00\x00\x00\x02\x00\x00\x00\x05\xc0\x53\xd7\xbb\x2a\x0d\x19\xc4\x40\x45\xd2\x9b\x79\x6d\xaa\x28\xc0\x53\xd7\xb5\xdb\x84\x1f\xb5\x40\x45\xd2\x9f\x26\xa1\x54\x79\xc0\x53\xd7\xb1\x20\x9e\xdb\xf9\x40\x45\xd2\xa1\xaf\x11\xd0\xe3\xc0\x53\xd7\xac\xf8\x86\x8e\xfb\x40\x45\xd2\xa4\x48\x49\x44\xed\xc0\x53\xd7\xab\xbf\x36\x0b\x55\x40\x45\xd2\xa5\x07\x8b\xe5\x7c\x00\x00\x00\x00\x02\x00\x00\x00\x02\xc0\x53\xd7\xab\xbf\x36\x0b\x55\x40\x45\xd2\xa5\x07\x8b\xe5\x7c\xc0\x53\xd7\xaa\xe5\x86\xd7\xf6\x40\x45\xd2\xa0\x9c\xc3\x19\xc6"),
			ndr: []byte("\x01\x07\x00\x00\x00\x03\x00\x00\x00\x01\x01\x00\x00\x00\x55\x0B\x36\xBF\xAB\xD7\x53\xC0\x7C\xE5\x8B\x07\xA5\xD2\x45\x40\x01\x02\x00\x00\x00\x05\x00\x00\x00\xC4\x19\x0D\x2A\xBB\xD7\x53\xC0\x28\xAA\x6D\x79\x9B\xD2\x45\x40\xB5\x1F\x84\xDB\xB5\xD7\x53\xC0\x79\x54\xA1\x26\x9F\xD2\x45\x40\xF9\xDB\x9E\x20\xB1\xD7\x53\xC0\xE3\xD0\x11\xAF\xA1\xD2\x45\x40\xFB\x8E\x86\xF8\xAC\xD7\x53\xC0\xED\x44\x49\x48\xA4\xD2\x45\x40\x55\x0B\x36\xBF\xAB\xD7\x53\xC0\x7C\xE5\x8B\x07\xA5\xD2\x45\x40\x01\x02\x00\x00\x00\x02\x00\x00\x00\x55\x0B\x36\xBF\xAB\xD7\x53\xC0\x7C\xE5\x8B\x07\xA5\xD2\x45\x40\xF6\xD7\x86\xE5\xAA\xD7\x53\xC0\xC6\x19\xC3\x9C\xA0\xD2\x45\x40"),
		},
	} {
		test(t, tc.g, tc.xdr, tc.ndr)
	}
}

func TestRandom(t *testing.T) {
	for _, tc := range testdata.Random {
		test(t, tc.G, nil, tc.WKB)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range testdata.Random {
			if _, err := Unmarshal(tc.WKB); err != nil {
				b.Errorf("unmarshal error %v", err)
			}
		}
	}
}

func BenchmarkMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range testdata.Random {
			if _, err := Marshal(tc.G, NDR); err != nil {
				b.Errorf("marshal error %v", err)
			}
		}
	}
}

func TestCrashes(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want error
	}{
		{
			s: "\x01\x03\x00\x00\x00\x04\x00\x00\x00\a\x00\x00tٽ&\xf2\xa6\xd0\x1a" +
				"\xce\xc7\x1a\xfd67\xa3\x98Y.\xa5\xfbH\x1b\xe7|\xbe\xac\xfd%" +
				";\x05\\\x90c\x83\xe9g\x01\xcbk\xa3\xc8\xdb\x0f\xae\x16bYl" +
				"\x1b\x1a\xae\xe0\x95=o\x85/\xec\xd2~\xf3\xce\xe7\xad\x04\x92\xc3\xea" +
				"r\xacE\xe3A\u008cR\x86sb\xd5sҙ\u007f\x82\xec\x88\xff" +
				"\x8aM\xa7\u007f;\x9b\x93\xa2tٽ&\xf2\xa6\xd0\x1a\xce\xc7\x1a\xfd" +
				"67\xa3\x98\x05\x00\x00\x004\xed\x19\x9c/\x8ej\ue643\x018" +
				"?\x01|\x02\xa2\xad\x18Wyʡ\xb4h\xc1j\xf6\xbb\xf0=\xbf" +
				"\x03d%\xe6PsyQ\xce4pѹ\x1dcR\xadr\x14\t" +
				"\x02pm\x86=_\xfb%\x81\"\xde\xdf4\xed\x19\x9c/\x8ej\xee" +
				"\x99\x83\x018?\x01|\x02\x05\x00\x00\x00\xfb#\xbf\xc8\xe2i\xe9'" +
				"<(\xa3\u05ccz\x06a\x8e\x17<\x956\xa4\\K\xccy\u05f7" +
				"\xcc\xdfԴp.\x9b\xce\xef0nx}\xe9\xfc\x10\xf7?\xc9\xcc" +
				"!,\xab\x15}*;\x84K\xeco\u07b6$_\xea\xfb#\xbf\xc8" +
				"\xe2i\xe9'<(\xa3\u05ccz\x06a\x04\x00\x00\x00\x8f\x8a\x9f9" +
				"\x81\x10h!N\xdcf\n\xf0-\xeaL\x02\xba\xe9\x03\xd6/G\xc2" +
				"\x1cj\r\xd8 \xbc\xd6r\x05աTS\xb3\xa5\xdc\xd8\xfb\")" +
				"\xab\x19\xf7̏\x8a\x9f9\x81\x10h!N\xdcf\n\xf0-\xeaL",
			want: wkbcommon.ErrGeometryTooLarge{Level: 1, N: 1946157063, Limit: wkbcommon.MaxGeometryElements[1]},
		},
	} {
		if _, err := Unmarshal([]byte(tc.s)); !reflect.DeepEqual(err, tc.want) {
			t.Errorf("Unmarshal([]byte(%#v)) == ..., %v, want %v", tc.s, err, tc.want)
		}
	}
}