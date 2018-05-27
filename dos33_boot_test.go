package main

import (
	"fmt"
	"testing"
	"time"

	"apple2/cpu"
	"apple2/disk"
	"apple2/keyboard"
	"apple2/mmu"
	"apple2/system"
	"apple2/utils"
	"apple2/video"
)

const dosDiskImage = "dos33.dsk"

func runDos33Boot(t *testing.T) {
	// Boot up DOS3.3
	utils.RunUntilBreakPoint(t, 0x0801, 2, false, "Boot0")
	utils.RunUntilBreakPoint(t, 0xb700, 2, false, "Boot1")
	utils.RunUntilBreakPoint(t, 0x9d84, 2, false, "Boot2")
	utils.RunUntilBreakPoint(t, 0xd7d2, 5, false, "JMP to basic interpreter NEWSTT")
}

func TestDOS33Boot(t *testing.T) {
	cpu.InitInstructionDecoder()
	mmu.InitRAM()
	mmu.InitApple2eROM()
	mmu.InitIO()
	disk.ReadDiskImage(dosDiskImage)
	cpu.Init()
	keyboard.Init()
	video.Init()
	system.Init()
	cpu.SetColdStartReset()
	cpu.Reset()

	t0 := time.Now()

	runDos33Boot(t)

	elapsed := float64(time.Since(t0) / time.Millisecond)
	fmt.Printf("CPU Cycles:    %d\n", system.FrameCycles)
	fmt.Printf("Time elapsed:  %0.2f ms\n", elapsed)
	fmt.Printf("Speed:         %0.2f cycles/ms\n", float64(system.FrameCycles)/elapsed)
}
