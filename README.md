# cool-ascii-faces for golang

golang implementation of [cool-ascii-faces](https://github.com/maxogden/cool-ascii-faces)

## Usage

```go
for c := range faces.Stream() {
	fmt.Println(c)
}
```

## Command-Line

```
$ files
(っ◕‿◕)っ
( ˘ ³˘)♥
~(˘▾˘~)
(ノ・∀・)ノ
(๑>ᴗ<๑)
┌( ಠ_ಠ)┘
(✌ﾟ∀ﾟ)☞
ヽ༼ຈل͜ຈ༽ง
ノ( ゜-゜ノ)
ᶘ ᵒᴥᵒᶅ
༼ ºل͟º༼
(｡◕‿‿◕｡)
(づ｡◕‿‿◕｡)づ
ॱ॰⋆(˶ॢ‾᷄﹃‾᷅˵ॢ)
(☞ﾟ∀ﾟ)☞
...
```

## Installation

```
$ go get github.com/mattn/cool-ascii-faces
```

## Author

Yasuhiro Matsumoto (a.k.a mattn)
