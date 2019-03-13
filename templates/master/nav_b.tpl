{{define "nav_b"}}
<nav class="navbar is-info" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">
      <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">
    </a>

    <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div class="navbar-menu" id="navbarBasicExample">
    <!-- <div class="navbar-start"> -->
    <div class="navbar-end">
      
      <!-- Admin -->
      {{if .Session }} {{if .Session.CheckPermission "admin"}}
      <div class="navbar-item has-dropdown is-hoverable">
        <a class="navbar-link">
          Admin
        </a>
        <div class="navbar-dropdown">
          <a class="navbar-item" href="/clean_sessions">
            Clean session
          </a>
<!--           <a class="navbar-item">
            Contact
          </a>
          <hr class="navbar-divider">
          <a class="navbar-item">
           Temp 
          </a> -->
        </div>
      </div>
      {{end}} {{end}}
      <!-- End Admin -->

      <!-- Student -->
      <div class="navbar-item has-dropdown is-hoverable">
        <a class="navbar-link" href="/stuent/all">
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
      <!-- End Student -->

      <!-- Login -->
      {{if not .Session}}<a class="navbar-item" href="/auth/signin"> Entrar </a>{{end}}
      <!-- User name -->
      {{if .Session}}<a class="navbar-item" href="/">{{.Session.UserName}}</a>{{end}}
      <!-- Logout -->
      {{if .Session}}<a class="navbar-item" href="/auth/signout"> Sair </a>{{end}}
    </div>

<!--     <div class="navbar-end">
      <div class="navbar-item">
        <div class="buttons">
          <a class="button is-primary">
            <strong>Sign up</strong>
          </a>
          <a class="button is-light">
            Log in
          </a>
        </div>
      </div>
    </div> -->

  </div>

</nav>
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
