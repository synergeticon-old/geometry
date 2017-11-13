package geometry

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/gonum/matrix/mat64"
	"github.com/google/uuid"
)

// PointCloud Represents an array of vectors
type PointCloud struct {
	Vectors []*mat64.Vector
}

// Add adds a Vector to Pointcloud
func (pC *PointCloud) Add(vec ...*mat64.Vector) {
	pC.Vectors = append(pC.Vectors, vec...)
}

// FillRandom fills pointcloud with random vectors
func (pC *PointCloud) FillRandom(count int) {
	for i := 0; i < count; i++ {
		vec := mat64.NewVector(3, []float64{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		})
		pC.Vectors = append(pC.Vectors, vec)
	}
}

// Length returns the amount of vertices in PoinCloud
func (pC *PointCloud) Length() int {
	return len(pC.Vectors)
}

// ReadPCD reads in PCD data from Point Cloud Library
func (pC *PointCloud) ReadPCD(path string) error {
	fileHandle, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	isPoint := false
	for fileScanner.Scan() {
		// line ist die jeweilige Zeile
		line := fileScanner.Text()

		if isPoint {
			// Die Punkte sind zeilenweise als "x y z" gespeichert und werden erstmal in ein unterarray gesplittet
			point := strings.Split(line, " ")

			// die unterarrays werden zu einem vektor zusammen geschrieben
			x, err := strconv.ParseFloat(point[0], 64)
			if err != nil {
				return err
			}
			y, err := strconv.ParseFloat(point[1], 64)
			if err != nil {
				return err
			}
			z, err := strconv.ParseFloat(point[2], 64)
			if err != nil {
				return err
			}

			vector := mat64.NewVector(3, []float64{
				x,
				y,
				z,
			})

			// der vector wird in die gesamte PointCloud ergänzt
			pC.Vectors = append(pC.Vectors, vector)
		}

		// Erst nach erkennen dieser Zeile werden Vectoren erstellt.
		if line == "DATA ascii" || line == "end_header" {
			isPoint = true
		}
	}
	return nil
}

// SavePLY saves a Pointcloud to PLY file
func (pC *PointCloud) SavePLY(path string) error {

	if pC.Vectors == nil {
		return errors.New("pointcloud is empty")
	}

	ply := []byte{}
	ply = append(ply, []byte("ply\n")...)
	ply = append(ply, []byte("format ascii 1.0\n")...)
	body := []byte{}

	var (
		x string
		y string
		z string
	)

	for _, vector := range pC.Vectors {
		x = strconv.FormatFloat(vector.At(0, 0), 'E', -1, 64)
		y = strconv.FormatFloat(vector.At(1, 0), 'E', -1, 64)
		z = strconv.FormatFloat(vector.At(2, 0), 'E', -1, 64)
		body = append(body, []byte(x+" "+y+" "+z+"\n")...)
	}

	ply = append(ply, []byte("element vertex "+strconv.Itoa(len(pC.Vectors))+"\n")...)
	ply = append(ply, []byte("property float x\n")...)
	ply = append(ply, []byte("property float y\n")...)
	ply = append(ply, []byte("property float z\n")...)
	ply = append(ply, []byte("end_header\n")...)
	ply = append(ply, body...)
	return ioutil.WriteFile(path, ply, 0644)

}

// Transform transforms pointcloud with transformation matrix
func (pC *PointCloud) Transform(transMat *TransMat) {
	for i, vector := range pC.Vectors {
		pC.Vectors[i] = transMat.Transform(vector)
	}
}

// ShowInMeshlab shows the Pointcloudobject in meshlab
func (pC *PointCloud) ShowInMeshlab() error {
	tmpPath := os.TempDir() + "/" + uuid.New().String() + ".ply"
	fmt.Println(tmpPath)
	err := pC.SavePLY(tmpPath)
	if err != nil {
		return err
	}

	meshlabExecPath := "meshlab"
	if runtime.GOOS == "windows" {
		meshlabExecPath = `C:\Program Files\VCG\MeshLab\meshlab.exe`
	}

	meshlab := exec.Command(meshlabExecPath, tmpPath)
	err = meshlab.Start()
	if err != nil {
		return errors.New("meshlab could not be started")
	}
	meshlab.Wait()
	return os.Remove(tmpPath)
}
