/**
 * Package indexing
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/10/15
 */

package indexing

import (
	"strconv"
	"testing"
)

type Scene int

const (
	SceneAlpha Scene = iota
	SceneBeta
	SceneGamma
	SceneAlphaLogin
	SceneAlphaCycle
	SceneDeltaAuth
	SceneDeltaAuthLogin
	SceneDeltaAuthCycle
	SceneCommon
)

type SceneMode int

const (
	SceneModeX SceneMode = iota + 1
	SceneModeY
	SceneModeLegacy
	SceneModeModern
	SceneModeAuthenticate
	SceneModeProcess
)

// IsValidSwitch is Switch-based IsValid implementation
func IsValidSwitch(scene Scene, mode SceneMode) bool {
	switch mode {
	case SceneModeX:
		switch scene {
		case SceneAlpha, SceneBeta, SceneGamma, SceneDeltaAuth:
			return true
		}
	case SceneModeY:
		switch scene {
		case SceneAlpha, SceneBeta, SceneGamma, SceneAlphaLogin, SceneAlphaCycle, SceneDeltaAuth, SceneDeltaAuthLogin, SceneDeltaAuthCycle:
			return true
		}
	case SceneModeLegacy:
		switch scene {
		case SceneAlpha, SceneBeta:
			return true
		}
	case SceneModeModern:
		switch scene {
		case SceneAlphaLogin, SceneAlphaCycle, SceneGamma, SceneDeltaAuthLogin, SceneDeltaAuthCycle:
			return true
		}
	case SceneModeAuthenticate:
		switch scene {
		case SceneAlphaLogin, SceneAlpha, SceneDeltaAuthLogin:
			return true
		}
	case SceneModeProcess:
		switch scene {
		case SceneAlphaCycle, SceneBeta, SceneDeltaAuthCycle:
			return true
		}
	default:
		return false
	}
	return false
}

// IsValidMap is Map-based IsValid implementation
func IsValidMap(scene Scene, mode SceneMode) bool {
	var scenes map[Scene]bool

	switch mode {
	case SceneModeX:
		scenes = map[Scene]bool{
			SceneAlpha:     true,
			SceneBeta:      true,
			SceneGamma:     true,
			SceneDeltaAuth: true,
		}
	case SceneModeY:
		scenes = map[Scene]bool{
			SceneAlpha:          true,
			SceneBeta:           true,
			SceneGamma:          true,
			SceneAlphaLogin:     true,
			SceneAlphaCycle:     true,
			SceneDeltaAuth:      true,
			SceneDeltaAuthLogin: true,
			SceneDeltaAuthCycle: true,
		}
	case SceneModeLegacy:
		scenes = map[Scene]bool{
			SceneAlpha: true,
			SceneBeta:  true,
		}
	case SceneModeModern:
		scenes = map[Scene]bool{
			SceneAlphaLogin:     true,
			SceneAlphaCycle:     true,
			SceneGamma:          true,
			SceneDeltaAuthLogin: true,
			SceneDeltaAuthCycle: true,
		}
	case SceneModeAuthenticate:
		scenes = map[Scene]bool{
			SceneAlphaLogin:     true,
			SceneAlpha:          true,
			SceneDeltaAuthLogin: true,
		}
	case SceneModeProcess:
		scenes = map[Scene]bool{
			SceneAlphaCycle:     true,
			SceneBeta:           true,
			SceneDeltaAuthCycle: true,
		}
	default:
		return false
	}

	_, isValid := scenes[scene]
	return isValid
}

// 准备测试用的场景和模式
var testScenes = []Scene{
	SceneAlpha,
	SceneBeta,
	SceneGamma,
	SceneAlphaLogin,
	SceneAlphaCycle,
	SceneDeltaAuth,
	SceneDeltaAuthLogin,
	SceneDeltaAuthCycle,
	SceneCommon,
}

var testModes = []SceneMode{
	SceneModeX,
	SceneModeY,
	SceneModeLegacy,
	SceneModeModern,
	SceneModeAuthenticate,
	SceneModeProcess,
}

// BenchmarkIsValidSwitch is benchmark for Switch-based IsValid
func BenchmarkIsValidSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 示例：循环遍历所有场景和模式
		for _, mode := range testModes {
			for _, scene := range testScenes {
				_ = IsValidSwitch(scene, mode)
			}
		}
	}
}

// BenchmarkIsValidMap is benchmark for Map-based IsValid
func BenchmarkIsValidMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 示例：循环遍历所有场景和模式
		for _, mode := range testModes {
			for _, scene := range testScenes {
				_ = IsValidMap(scene, mode)
			}
		}
	}
}

func BenchmarkIsValidSwitchRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sceneStr := strconv.Itoa(i % len(testScenes))
		modeStr := strconv.Itoa(i % len(testModes))
		scene, _ := strconv.Atoi(sceneStr)
		mode, _ := strconv.Atoi(modeStr)
		_ = IsValidSwitch(Scene(scene), SceneMode(mode))
	}
}

func BenchmarkIsValidMapRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sceneStr := strconv.Itoa(i % len(testScenes))
		modeStr := strconv.Itoa(i % len(testModes))
		scene, _ := strconv.Atoi(sceneStr)
		mode, _ := strconv.Atoi(modeStr)
		_ = IsValidMap(Scene(scene), SceneMode(mode))
	}
}
