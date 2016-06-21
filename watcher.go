package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func cleanPath(path string) string {
	path, _ = filepath.Abs(path)
	filepath.Clean(path)
	return filepath.ToSlash(path)

}

func watchFolder(path string) {
	watcher, err := fsnotify.NewWatcher()
	lastModified := time.Now()
	checkError(err)

	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				current := time.Now()
				ignored := inSlice(options.IgnoredFolders, func(dir string) bool {
					return strings.HasPrefix(cleanPath(ev.Name), cleanPath(dir))
				})
				if ignored {
					break
				}
				ignored = inSlice(options.IgnoredExtensions, func(name string) bool {
					return name == filepath.Ext(ev.Name)
				})

				if ignored {
					break
				}

				if current.Sub(lastModified) > options.Delay*time.Millisecond {
					msg := ev.String()
					i := strings.LastIndex(msg, ":")
					//FORMAT -> "filename":OP
					//cremove quotes and call cleanPath for filename
					if i != -1 {
						watchLog("\n\nReceived %s", cleanPath(msg[1:i-1])+":"+msg[i+1:])
					} else {
						watchLog("\n\nReceived %s", ev.String()) //should never happen
					}
					event <- ev.String()
					lastModified = current
				} else {
					//	watchLog("Discarding %s", ev.String())
				}

			case err := <-watcher.Errors:
				if err != nil {
					watchLog("Error: %s", err.Error())
				}

			}
		}

	}()

	err = watcher.Add(path)

	checkError(err)

}

func watch(path string) {
	watchLog("\n\n")
	filepath.Walk(options.Root, func(path string, info os.FileInfo, err error) error {
		//ignore error as path will always be valid here
		fileInfo, _ := os.Stat(path)
		if !fileInfo.IsDir() {
			return nil
		}
		path = cleanPath(path)
		ignored := inSlice(options.IgnoredFolders, func(dir string) bool {
			return strings.HasPrefix(path, cleanPath(dir))

		})

		if !ignored {
			watchLog("Watching %s", path)
			watchFolder(path)
		} else {
			watchLog("Ignoring %s", path)
		}

		return nil
	})

}
