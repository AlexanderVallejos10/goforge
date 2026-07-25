<script setup>
import { ref } from "vue"

import {
  Add,
  Commit,
  GetBranch,
  GetCoreInfo,
  GetStatus
} from "../wailsjs/go/main/App"


const coreMensaje = ref("")

const rama = ref("")

const cambios = ref([])

const mensajeCommit = ref("")

const repositorio = ref(
  "C:/temp/demo-goforge"
)



function conectarCore(){

  GetCoreInfo()
    .then(resultado => {

      coreMensaje.value = resultado

    })

}



function cargarRepositorio(){

  GetBranch(
    repositorio.value
  )
  .then(resultado => {

    rama.value = resultado

  })


  GetStatus(
    repositorio.value
  )
  .then(resultado => {

    cambios.value = resultado

  })

}



function agregarArchivos(){

  Add(
    repositorio.value
  )
  .then(resultado => {

    coreMensaje.value = resultado

    cargarRepositorio()

  })

}



function crearCommit(){

  if(
    mensajeCommit.value.trim() === ""
  ){

    coreMensaje.value =
      "Escribe un mensaje de commit"

    return

  }


  Commit(
    repositorio.value,
    mensajeCommit.value
  )
  .then(resultado => {

    coreMensaje.value = resultado

    mensajeCommit.value = ""

    cargarRepositorio()

  })

}

</script>



<template>

<div class="app">


<header class="topbar">

<h1>
goGitDesktop
</h1>


<span class="version">
GoForge Core v0.1
</span>


</header>



<main class="layout">


<aside class="sidebar">


<h3>
Repositorios
</h3>


<div class="repo">

📁 demo-goforge

</div>



<button
@click="cargarRepositorio"
>

Abrir repositorio

</button>



</aside>




<section class="content">



<div class="panel">


<h2>
Estado Core
</h2>


<button
class="primary"
@click="conectarCore"
>

Conectar Core

</button>


<p class="success">

{{ coreMensaje }}

</p>


</div>





<div class="panel">


<h2>
Repositorio
</h2>


<p>

Ruta:

{{ repositorio }}

</p>


<p>

Rama actual:

<strong>
{{ rama }}
</strong>

</p>


</div>





<div class="panel">


<h2>
Cambios
</h2>



<div
v-if="cambios.length === 0"
class="empty"
>

No hay cambios

</div>



<div
v-for="(cambio,index) in cambios"
:key="index"
class="item"
>


<span>

{{ cambio.archivo }}

</span>


<strong>

{{ cambio.tipo }}

</strong>


</div>



</div>





<div class="panel">


<h2>
Acciones
</h2>



<input
v-model="mensajeCommit"
placeholder="Mensaje del commit"
/>



<div class="actions">



<button
@click="agregarArchivos"
>

Add

</button>




<button
@click="crearCommit"
>

Commit

</button>



<button>
Branch
</button>



<button>
Checkout
</button>



<button>
Diff
</button>



<button>
Restore
</button>



<button>
Reset
</button>



</div>



</div>



</section>



</main>



</div>


</template>




<style scoped>


.app {

height:100vh;

background:#0d1117;

color:#e6edf3;

font-family:
Segoe UI,
sans-serif;

}



.topbar {

height:60px;

display:flex;

align-items:center;

padding:0 25px;

background:#161b22;

border-bottom:
1px solid #30363d;

}



.topbar h1 {

color:#58a6ff;

font-size:32px;

}



.version {

margin-left:auto;

color:#8b949e;

}



.layout {

display:flex;

height:calc(100vh - 60px);

}



.sidebar {

width:260px;

background:#161b22;

padding:20px;

border-right:
1px solid #30363d;

}



.sidebar button {

width:100%;

padding:12px;

margin-top:20px;

background:#238636;

border:none;

color:white;

border-radius:6px;

cursor:pointer;

}



.repo {

background:#21262d;

padding:15px;

border-radius:8px;

margin-top:20px;

}



.content {

flex:1;

padding:25px;

overflow:auto;

}



.panel {

background:#161b22;

border:
1px solid #30363d;

border-radius:10px;

padding:20px;

margin-bottom:20px;

}



.panel h2 {

color:#58a6ff;

}



.primary {

background:#1f6feb;

border:none;

color:white;

padding:10px 20px;

border-radius:6px;

cursor:pointer;

}



.success {

color:#3fb950;

margin-top:15px;

}



.item {

display:flex;

justify-content:space-between;

background:#21262d;

padding:12px;

border-radius:6px;

margin-top:10px;

}



.item strong {

color:#f0883e;

}



.empty {

color:#8b949e;

}



input {

width:300px;

padding:10px;

margin-bottom:15px;

background:#21262d;

color:white;

border:1px solid #30363d;

border-radius:6px;

}



.actions {

display:flex;

gap:10px;

flex-wrap:wrap;

}



.actions button {

background:#1f6feb;

color:white;

border:none;

padding:10px 20px;

border-radius:6px;

cursor:pointer;

}


</style>