package main

import (
	"html/template"
	"net/http"
)

// Tipe data untuk menyimpan konten syarat dan ketentuan
type TermsAndConditions struct {
	Content template.HTML
}

func main() {
	http.HandleFunc("/", serveTermsAndConditions)
	http.ListenAndServe(":8080", nil)
}

func serveTermsAndConditions(w http.ResponseWriter, r *http.Request) {
	// Contoh konten syarat dan ketentuan (dapat diganti sesuai kebutuhan)
	content := `
		<h1>Terms & Conditions</h1>
		<p>Silakan baca dan pahami syarat dan ketentuan penggunaan kami sebelum menggunakan layanan kami.</p>
        <h3>Lorem ipsum</h3>
        <p>
          Lorem, ipsum dolor sit amet consectetur adipisicing elit. Odio temporibus,
          id alias error iusto eos deleniti, voluptatibus reprehenderit repellat
          tenetur minus iure voluptatem nesciunt assumenda dolor tempora, facilis
          minima magnam totam ullam. Pariatur, possimus, fugit autem sapiente
          aperiam a vitae temporibus consequuntur ipsum voluptas, adipisci officiis
          nam nihil explicabo vero repudiandae odio quisquam harum. Incidunt quo
          quaerat earum corrupti perspiciatis amet enim nobis itaque hic tempore.
          Harum eius consequatur atque neque minima repellat eligendi, nostrum
          cumque quo dolore. Iste voluptatum temporibus nemo repellat? Praesentium
          facilis quam dolorum non dignissimos asperiores est tempore quas, tenetur
          consequuntur quaerat molestiae assumenda! Accusamus, rerum!
        </p>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Quibusdam, fugit.
        </p>
        <h3>Lorem ipsum dolor sit amet.</h3>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequatur
          reprehenderit voluptates ab, deleniti nemo autem suscipit iusto
          accusantium ratione, optio unde molestiae voluptatum? Aut neque quaerat
          dignissimos reprehenderit, quidem tempore, cum optio similique iusto,
          dolores molestias veritatis rerum suscipit dolore molestiae laboriosam
          inventore corporis? Rerum, quisquam facere esse itaque rem dicta pariatur,
          similique et unde voluptatum est, magni inventore ex eos? Temporibus,
          suscipit accusamus ut laborum molestias similique voluptate ullam porro
          deserunt ex veritatis vitae architecto voluptas necessitatibus rerum
          quidem dolorum! Temporibus quae molestias deleniti fugit officiis ipsa ut.
          Nemo alias obcaecati consequatur fuga id ipsam ab beatae omnis odio.
        </p>
		<!-- ... Isi syarat dan ketentuan di sini ... -->
	`

	terms := TermsAndConditions{
		Content: template.HTML(content),
	}

	tmpl, err := template.New("terms").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Syarat dan Ketentuan</title>
		</head>
		<body>
			{{.Content}}
		</body>
		</html>
	`)

	if err != nil {
		http.Error(w, "Terjadi kesalahan internal", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, terms)
	if err != nil {
		http.Error(w, "Terjadi kesalahan internal", http.StatusInternalServerError)
		return
	}
}
