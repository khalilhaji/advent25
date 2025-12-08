package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance_happy(t *testing.T) {
	assert.Equal(t, 3., distance(box{3, 0, 0, 0}, box{0, 0, 0, 0}))
}

func TestComputeDirectConnections_happy(t *testing.T) {
	boxes, err := parse([]string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	})

	assert.NoError(t, err)

	connections := computeDirectConnections(boxes)
	assert.Equal(t, directConnection{
		distance(boxes[0], boxes[19]),
		0,
		19,
	}, connections[0])
	assert.Equal(t, directConnection{
		distance(boxes[0], boxes[7]),
		0,
		7,
	}, connections[1])
	assert.Equal(t, directConnection{
		distance(boxes[2], boxes[13]),
		2,
		13,
	}, connections[2])
	assert.Equal(t, directConnection{
		distance(boxes[7], boxes[19]),
		7,
		19,
	}, connections[3])
}

func TestMultiplyLongestCircuits_happy(t *testing.T) {
	boxes, err := parse([]string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	})

	assert.NoError(t, err)

	connections := computeDirectConnections(boxes)
	err = makeConnections(10, boxes, connections)
	assert.NoError(t, err)

	res := multiplyLongestCircuits(boxes)
	assert.Equal(t, 40, res)
}

func Test_makeConnectionsV2(t *testing.T) {
	boxes, err := parse([]string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	})

	assert.NoError(t, err)

	connections := computeDirectConnections(boxes)
	res := makeConnectionsV2(boxes, connections)
	assert.Equal(t, 25272, res)
}
