run: skills.json education.json references.json experience.json resume.go
	@echo "Creating README.md..."
	go run resume.go

clean:
	        @echo "Cleaning up..."
	        rm README.md
