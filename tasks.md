# Implementation Tasks

- [x] Initialize Go module (`go mod init github.com/zurustar/ani`) <!-- id: 10 -->
- [x] Install dependencies (`github.com/fogleman/gg`) <!-- id: 11 -->
- [x] Create `animator` package struct and logic <!-- id: 12 -->
- [x] Create `renderer` package with `gg` and transparent background logic <!-- id: 13 -->
- [x] Implement `main.go` CLI with flags <!-- id: 14 -->
- [x] Verify Output (Manual visual check of generated GIF) <!-- id: 15 -->
- [x] TDD: Create Tests & Refactor for Testability <!-- id: 21 -->
    - [x] Create `renderer` tests (Color Parsing) <!-- id: 22 -->
    - [x] Create `animator` tests (Frame/Step Calculation) <!-- id: 23 -->
- [x] Implement Background Color Support <!-- id: 16 -->
    - [x] Update `renderer` to support custom background color <!-- id: 17 -->
    - [x] Update `animator` to handle background color in palette generation <!-- id: 18 -->
    - [x] Update `main.go` to parse `-bg` flag (strict hex) <!-- id: 19 -->
    - [x] Verify background color feature <!-- id: 20 -->
