// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package markdowntohtml

import (
	"github.com/andreaskoch/allmark2/common/logger"
	"github.com/andreaskoch/allmark2/common/paths"
	"github.com/andreaskoch/allmark2/model"
	"github.com/andreaskoch/allmark2/services/conversion/markdowntohtml/audio"
	"github.com/andreaskoch/allmark2/services/conversion/markdowntohtml/files"
	"github.com/andreaskoch/allmark2/services/conversion/markdowntohtml/markdown"
)

type Converter struct {
	logger logger.Logger
}

func New(logger logger.Logger) (*Converter, error) {
	return &Converter{
		logger: logger,
	}, nil
}

// Convert the supplied item with all paths relative to the supplied base route
func (converter *Converter) Convert(pathProvider paths.Pather, item *model.Item) (convertedContent string, conversionError error) {

	converter.logger.Debug("Converting item %q.", item)

	content := item.Content

	// markdown extension: audio
	audioConverter := audio.New(pathProvider, item.Files())
	content, audioConversionError := audioConverter.Convert(content)
	if audioConversionError != nil {
		converter.logger.Warn("Error while converting audio extensions. Error: %s", audioConversionError)
	}

	// markdown extension: files
	filesConverter := files.New(pathProvider, item.Files())
	content, filesConversionError := filesConverter.Convert(content)
	if filesConversionError != nil {
		converter.logger.Warn("Error while converting files extensions. Error: %s", filesConversionError)
	}

	// markdown to html
	content = markdown.Convert(content)

	return content, nil
}
