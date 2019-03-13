{{define "header"}}
<style type="text/css">
/*  .navbar {
    min-height: 4em;
  }*/
</style>
<header>
  <nav class="navbar">
    <div class="container">

      <div class="navbar-brand">
        <a class="navbar-item" href="/">
          <img style="max-height: 36px;" src="/static/img/sail2.jpg">
          <p>&nbspVelejar!</p>
        </a>
        <a role="button" class="navbar-burger burger" data-target="mainNav">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div class="navbar-menu" id="mainNav">

        <div class="navbar-start">
          <!-- Institutional. -->
          <a class="navbar-item" href="/info/institutional"> Institucional </a>
          <!-- Admin -->
          <div class="navbar-item has-dropdown is-hoverable">
            <!-- <a class="navbar-link is-arrowless"> -->
            <a class="navbar-link">
              Cursos e serviços
            </a>
            <div class="navbar-dropdown">
              <a class="navbar-item" href="/info">
                Aulas de vela para crianças
              </a>
              <a class="navbar-item" href="/info">
                Aulas de vela para adultos
              </a>
              <a class="navbar-item" href="/info">
                Passeios
              </a>
              <a class="navbar-item" href="/info">
                Aluguel de veleiros
              </a>
              <a class="navbar-item" href="/info">
                Aulas de remo
              </a>
              <a class="navbar-item" href="/info">
                Aluguel de caiaques e bikes aquáticas
              </a>
              <a class="navbar-item" href="/info">
                Fábrica náutica
              </a>
            </div>
          </div>
          <a class="navbar-item" href="/info">Projetos e iniciativas</a>
          <a class="navbar-item" href="/info">Contato</a>
          <a class="navbar-item" href="/info">Blog</a>
          <a class="navbar-item" href="/info">Área do aluno</a>
        </div>

        <!-- <div class="navbar-start"> -->
        <div class="navbar-end">
          <!-- Admin -->
          {{if .Session }} {{if .Session.CheckPermission "admin"}}
          <div class="navbar-item has-dropdown is-hoverable">
            <!-- <a class="navbar-link is-arrowless"> -->
            <a class="navbar-link">
              Admin
            </a>
            <div class="navbar-dropdown">
              <a class="navbar-item" href="/clean_sessions">
                Clean session
              </a>
            </div>
          </div>
          {{end}} {{end}}
          <!-- End Admin -->
          <!-- Student -->
          {{if .Session }} {{if .Session.CheckPermission "editStudent"}}
          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link">
              Aluno
            </a>
            <div class="navbar-dropdown">
              <a class="navbar-item" href="/student/all">
                Alunos
              </a>
              <a class="navbar-item" href="/student/new">
                Novo aluno
              </a>
            </div>
          </div>
          {{end}} {{end}}
          <!-- End Student -->
          <!-- Login -->
          {{if not .Session}}<a class="navbar-item" href="/auth/signin"> Entrar </a>{{end}}
          <!-- User name -->
          {{if .Session}}<a class="navbar-item" href="/">{{.Session.UserName}}</a>{{end}}
          <!-- Logout -->
          {{if .Session}}<a class="navbar-item" href="/auth/signout"> Sair </a>{{end}}
        </div>
      </div>
    </div>
  </nav>
</header>
<script type="text/javascript">
  document.addEventListener('DOMContentLoaded', () => {
    // Get all "navbar-burger" elements
    const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);
    // Check if there are any navbar burgers
    if ($navbarBurgers.length > 0) {
      // Add a click event on each of them
      $navbarBurgers.forEach( el => {
        el.addEventListener('click', () => {
          // Get the target from the "data-target" attribute
          const target = el.dataset.target;
          const $target = document.getElementById(target);
          // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
          el.classList.toggle('is-active');
          $target.classList.toggle('is-active');
        });
      });
    }
  });
</script>
{{end}}
