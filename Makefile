TARGET := vulcan

build: export

compile:
	@echo "Compiling the project..."
	@go build -tags embed -o bin/$(TARGET) main.go
	@echo "Project compiled successfully!"

export: compile
	@echo "\nExporting the project into the system..."
	@cp bin/$(TARGET) /usr/local/bin
	@echo "Project exported successfully!"

clean:
	@echo "Cleaning the project..."
	@rm -rf bin/*
	@rm -rf usr/local/$(TARGET)
	@echo "Project cleaned successfully!"
