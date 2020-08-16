```
OsmTagFilter -input-format pbf -input croatia-latest.osm.pbf \
 -key place -value village \
 -output-format pbf -output croatia-villages.pbf

```

```
OsmTagFilter -input-format pbf -input croatia-latest.osm.pbf \
 -key place -value town \
 -output-format pbf -output croatia-towns.pbf

```

```
OsmKeyFilter -input-format pbf -input croatia-latest.osm.pbf \
    -key highway \
    -output-format pbf -output ways.pbf
```
