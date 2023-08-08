<script>
  import {SaveTasks, LoadTasks} from '../wailsjs/go/main/App.js'
  import {onMount} from 'svelte'
  import {fly} from 'svelte/transition'

  let inputText = "";
  let tasks = [];
  let error = "";
  let index = 0;

  function addTask() {
    if (inputText === "") {
      error = "• Enter something...";
      return;
    }
    tasks = [{
      Id: Date.now(),
      Desc: inputText
    }].concat(tasks);
    inputText = "";
    saveAll()
  }

  function removeTask(el_id) {
    tasks = tasks.filter(el => el.Id !== el_id);
    saveAll();
  }

  function saveAll() {
    SaveTasks(tasks.map(el => el.Id+": "+el.Desc).join("\n"));
  }

  function addTaskCustomArgs(idx, desc) {
    tasks = tasks.concat([{
      Id: idx,
      Desc: desc
    }]);
  }

  function loadAll() {
    LoadTasks().then(tasks => {
      if (tasks === "") {
        return;
      }
      for (const task_line of tasks.split("\n")) {
        addTaskCustomArgs(Number(task_line.split(": ")[0]), task_line.split(": ").slice(1).join(": "));
      }
    });
  }

  onMount(() => {
    loadAll();
  });
</script>

<main>
  <form on:submit|preventDefault={()=>{addTask();index=0;}}>
    <input type="text" bind:value={inputText} on:keyup={ev => {
      if (ev.code.toString() == "ArrowUp") {
        inputText = tasks[index].Desc;
        if (inputText !== "") {
          error = "";
        }
        index++;
        if (index>tasks.length-1) {
          index = 0;
        }
      }
    }} on:input={() => {error="";index=0;}} placeholder="What do you like do?">
    <p class="red-color error">{error}</p>
  </form>
  {#if tasks.length > 1}
    <button class="clearall" on:click={() => {
        if (confirm("All your tasks are deleted. Do you agree?")) {
          tasks = [];
          saveAll();
        }
      }}>Clear all</button>
  {/if}
  <div class="tasks">
  {#each tasks as task (task.Id)}
    <div class="task" transition:fly={{y: 10, duration: 500}}>
      <span>
        <b>Task:</b> 
        <span title="{task.Desc}">{task.Desc}</span>
      </span>
      <button on:click={() => {removeTask(task.Id);}}>remove</button>
    </div>
  {:else}
    <p class="nothinghere">Nothing here!</p>
  {/each}
  </div>
</main>