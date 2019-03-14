{{template "base_p" .}}
{{define "title"}}Aulas de vela para crianças{{end}}
{{define "main"}}
<style type="text/css">
    body {
        color: #777;
    }
    #layout {
        position: relative;
        left: 0;
        padding-left: 0;
    }
        #layout.active #menu {
            left: 150px;
            width: 150px;
        }

        #layout.active .menu-link {
            left: 150px;
        }

    .header {
        margin: 0;
        padding-top: 2.5em;
        text-align: center;
        border-bottom: 1px solid #eee;
    }
    .header h1, .header h2 {
        font-weight: 500;
    }
    .header h1 {
        font-size: 3em;
        color: #333;
        margin: .2em 0;
    }
    .header h2 {
        color: #ccc;
        margin-top: 0;
    }
    .content {
        max-width: 800px;
        margin: 0 auto;
        padding: 0 2em;
        margin-bottom: 50px;
        line-height: 1.6em;
    }
    .content h2 {
        color: #888;
        font-weight: 500;
        margin: 50px 0 20px 0;
    }
    #menu {
        margin-left: 0px; /* "#menu" width */
        width: 150px;
        position: fixed;
        top: 0;
        left: 0;
        bottom: 0;
        z-index: 1000; /* so the menu or its navicon stays above all content */
        background: #191818;    /* ? */
        overflow-y: auto;   /* ? */
        -webkit-overflow-scrolling: touch;
    }
    /*
    All anchors inside the menu should be styled like this.
    */
    #menu a {
        color: #999;
        border: none;
        padding: 0.6em 0 0.6em 0.6em;
    }
    /*
    Remove all background/borders, since we are applying them to #menu.
    */
     #menu .pure-menu,
     #menu .pure-menu ul {
        border: none;
        background: transparent;
    }

        /*
    Add that light border to separate items into groups.
    */
    #menu .pure-menu ul,
    #menu .pure-menu .menu-item-divided {
        border-top: 1px solid #333;
    }
        /*
        Change color of the anchor links on hover/focus.
        */
        #menu .pure-menu li a:hover,
        #menu .pure-menu li a:focus {
            background: #333;
        }

        /*
    This styles the selected menu item `<li>`.
    */
    #menu .pure-menu-selected,
    #menu .pure-menu-heading {
        background: #1f8dd6;
    }
        /*
        This styles a link within a selected menu item `<li>`.
        */
        #menu .pure-menu-selected a {
            color: #fff;
        }

    /*
    This styles the menu heading.
    */
    #menu .pure-menu-heading {
        font-size: 110%;
        color: #fff;
        margin: 0;
    }


    .menu-link {
        position: fixed;
        display: block; /* show this only on small screens */
        top: 0;
        left: 0; /* "#menu width" */
        background: #000;
        background: rgba(0,0,0,0.7);
        font-size: 10px; /* change this value to increase/decrease button size */
        z-index: 10;
        width: 2em;
        height: auto;
        padding: 2.1em 1.6em;
    }
        .menu-link:hover,
        .menu-link:focus {
            background: #000;
        }

        .menu-link span {
            position: relative;
            display: block;
        }

        .menu-link span,
        .menu-link span:before,
        .menu-link span:after {
            background-color: #fff;
            width: 100%;
            height: 0.2em;
        }

            .menu-link span:before,
            .menu-link span:after {
                position: absolute;
                margin-top: -0.6em;
                content: " ";
            }

            .menu-link span:after {
                margin-top: 0.6em;
            }

    /* -- Responsive Styles (Media Queries) ------------------------------------- */

    /*
    Hides the menu at `48em`, but modify this based on your app's needs.
    */
    @media (min-width: 48em) {

        .header,
        .content {
            padding-left: 2em;
            padding-right: 2em;
        }

        #layout {
            padding-left: 150px; /* left col width "#menu" */
            left: 0;
        }
        #menu {
            left: 150px;
        }

        .menu-link {
            position: fixed;
            left: 150px;
            display: none;
        }

        #layout.active .menu-link {
            left: 150px;
        }
    }

    @media (max-width: 48em) {
        /* Only apply this when the window is small. Otherwise, the following
        case results in extra padding on the left:
            * Make the window small.
            * Tap the menu to trigger the active state.
            * Make the window large again.
        */
        #layout.active {
            position: relative;
            left: 150px;
        }
    }

</style>
<div id="layout">
    <!-- Menu toggle -->
    <a href="#menu" id="menuLink" class="menu-link">
        <!-- Hamburger icon -->
        <span></span>
    </a>
    <div id="menu">
        <div class="pure-menu">
            <a class="pure-menu-heading" href="#">Company</a>
            <ul class="pure-menu-list">
                <li class="pure-menu-item"><a href="#" class="pure-menu-link">Home</a></li>
                <li class="pure-menu-item"><a href="#" class="pure-menu-link">About</a></li>

                <li class="pure-menu-item menu-item-divided pure-menu-selected">
                    <a href="#" class="pure-menu-link">Services</a>
                </li>

                <li class="pure-menu-item"><a href="#" class="pure-menu-link">Contact</a></li>
            </ul>
        </div>
    </div>
    <div id="main">
        <div class="header">
            <h1>Escolinha de Optimist</h1>
            <h2>Um barco para crianças e jovens</h2>
        </div>
        <div class="content">
            <h2>O barco</h2>

            <p> Criança também veleja, e os pequenos tem uma classe de vela só para eles: a Optimist! </p>

            <p> 
                Ensinamos a arte de velejar de forma lúdica, trabalhando o respeito ao meio ambiente, 
                a competição saudável e o companheirismo.
                Brincando, as crianças aprendem a reconhecer as partes do barco, a observar o vento
                e a definir estratégias para velejar.
            </p>

            <h2>Benefícios da prática de vela</h2>
            <ul>
                <li>Melhora a autoconfiança</li>
                <li>Trabalha a disciplina e concentração</li>
                <li>Estimula a responsabilidade</li>
                <li>Desenvolve o trabalho em equipe</li>
                <li>Incentiva o respeito ao meio ambiente</li>
                <li>Trabalha a concentração e tomada de decisão</li>
                <li>Desenvolve a coordenação motora, o senso de direção e a noção espacial</li>
            </ul>
          
            <p>Agende uma aula experimental e conte conosco!</p>
        </div>
    </div>
</div>

{{end}}