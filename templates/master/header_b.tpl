{{define "header_b"}}
<header class="container mt-3">
    <!-- <nav class="navbar navbar-expand-lg navbar-light" style="background-color: #e3f2fd;"> -->
    <!-- <nav class="navbar navbar-expand-lg navbar-light bg-white"> -->
    <!-- <nav class="navbar navbar-expand-lg navbar-light bg-light"> -->
    <nav class="navbar navbar-expand-lg navbar-light bg-white">
        <a class="navbar-brand" href="/_b">
            <img src="/static/img/sail2.jpg" width="30" height="30" alt="">
        </a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item"><a class="nav-link" href="/info/institutional">Institucional</a></li>
                <li class="nav-item dropdown">
                    <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        Cursos e serviços   
                    </a>
                    <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                        <a class="dropdown-item" href="/info">Aulas de vela para crianças</a>
                        <a class="dropdown-item" href="/info">Aulas de vela para adultos</a>
                        <a class="dropdown-item" href="/info">Aulas de remo</a>
                        <div class="dropdown-divider"></div>
                        <a class="dropdown-item" href="/info">Aluguel de veleiros</a>
                        <a class="dropdown-item" href="/info">Aluguel de caiaques e bikes aquáticas</a>
                        <div class="dropdown-divider"></div>
                        <a class="dropdown-item" href="/info">Passeios</a>
                    </div>
                </li>
                <li class="nav-item"><a class="nav-link" href="/info">Projetos e iniciativas</a></li>
                <li class="nav-item"><a class="nav-link" href="/info">Contato</a></li>
                <li class="nav-item"><a class="nav-link" href="/info">Blog</a></li>
                <li class="nav-item"><a class="nav-link" href="/info">Área do aluno</a></li>
            </ul>
            <!-- <ul class="navbar-nav mr-auto"> -->
            <ul class="navbar-nav">
                {{if .Session }} {{if .Session.CheckPermission "admin"}}
                <li class="nav-item dropdown">
                    <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        Admin
                    </a>
                    <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                        <a class="dropdown-item" href="/clean_sessions">Clean sessions</a>
                  </div>
                </li>
                {{end}}{{end}}
                {{if .Session }} {{if .Session.CheckPermission "editStudent"}}
                <li class="nav-item dropdown">
                    <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        Aluno
                    </a>
                    <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                        <a class="dropdown-item" href="/student/all">Alunos</a>
                        <a class="dropdown-item" href="/student/new">Novo aluno</a>
                    </div>
                </li>
                {{end}}{{end}}
                {{if not .Session}}<li class="nav-item"><a class="nav-link" href="/auth/signin">Entrar</a></li>{{end}}
                {{if .Session}}<li class="nav-item"><a class="nav-link" href="/">{{.Session.UserName}}o</a></li>{{end}}
                {{if .Session}}<li class="nav-item"><a class="nav-link" href="/auth/signout">Sair</a></li>{{end}}
            </ul>
        </div>
    </nav>
</header>
{{end}}
