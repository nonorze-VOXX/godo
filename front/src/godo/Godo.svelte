<script lang="ts">
    const myUrl = "http://127.0.0.1:4081/godoList";

    let Stime = "2018-12-10T13:45:00.000Z";
    let Etime = "2018-12-10T13:45:00.000Z";
    let Done = true;
    let Title = "godo Title";
    let result = null;

    async function postGodo() {
        if (Title != "") {
            const res = await fetch(myUrl, {
                method: "POST",
                body: JSON.stringify({
                    Title,
                    Stime,
                    Etime,
                    Done,
                }),
                headers: {
                    "Content-Type": "application/json",
                    // 'Content-Type': 'application/x-www-form-urlencoded',
                },
            });
            const json = await res.json();
            result = JSON.stringify(json);
        } else {
            result = "title is no fill";
        }
    }

    import GetGodoList from "./GetGodoList.svelte";
</script>

<main>
    <legend>submit new godo</legend>
    <input bind:value={Title} />
    <button type="button" on:click|preventDefault={postGodo}>new</button>
    <pre>
        {result}
    </pre>
    <GetGodoList />
</main>

<style>
</style>
