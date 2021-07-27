const btnGit = document.querySelector('.git-button');
const gitText = document.querySelector('.git-text');
const uninstall = document.querySelector('.uninstall');
const plug = document.querySelector('.porkPlugger');
const unplug = document.querySelector('.porkUnplugger');

btnGit.addEventListener('click', async () => {
	gitText.innerText = 'Installing powercord...';
	gitText.innerText = await installPC();
});

uninstall.addEventListener('click', async () => {
	gitText.innerText = 'Uninstalling powercord...';
	gitText.innerText = await uninstallPC();
});

plug.addEventListener('click', async () => {
	gitText.innerText = 'Plugging powercord...';
	gitText.innerText = await plugPowercord();
});

unplug.addEventListener('click', async () => {
	gitText.innerText = 'Unplugging powercord...';
	gitText.innerText = await unplugPowercord();
});

async function setup() {
	const cando = JSON.parse(await canDo());
	if (!cando.ok) {
		btnGit.disabled = true;
		uninstall.disabled = true;
		plug.disabled = true;
		unplug.disabled = true;
		gitText.innerHTML = `<a target=_blank href=${cando.link}>${cando.text}<a/>`;
	}
}

window.onload = setup;
