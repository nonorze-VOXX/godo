<script lang="ts">
    const myUrl = "http://127.0.0.1:4081/godoList";

    let Stime = "2018-12-10T13:45:00.000Z";
    let Etime = "2018-12-10T13:45:00.000Z";
    let Done = true;
    let Title = "godo Title";
    let result = null;

    async function postGodo() {
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
    }
</script>

<main>
    <form>
        <fieldset>
            <legend>submit new godo</legend>
            <input bind:value={Title} />
            <button type="button" on:click|preventDefault={postGodo} />
        </fieldset>
    </form>
    <pre>
        {result}
    </pre>
</main>

<style>
</style>
