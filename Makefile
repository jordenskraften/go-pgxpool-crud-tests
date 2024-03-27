mup:
	cd ./migrations &&  goose postgres "user=postgres password=mysecretpassword host=localhost dbname=interview_spellbook sslmode=disable" up

mdown:
	cd ./migrations && goose postgres "user=postgres password=mysecretpassword host=localhost dbname=interview_spellbook sslmode=disable" up