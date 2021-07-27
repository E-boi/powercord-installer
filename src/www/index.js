const btnGit = document.querySelector('.git-button');
const gitText = document.querySelector('.git-text');
const uninstall = document.querySelector('.uninstall');
const plug = document.querySelector('.porkPlugger');
const unplug = document.querySelector('.porkUnplugger');
const btnTheme = document.querySelector('.themeDownloader');
const btnPlugin = document.querySelector('.pluginDownloader');
const divPlugins = document.querySelector('.plugins');
let installed = false;

btnGit.addEventListener('click', async () => {
	gitText.innerText = 'Installing powercord...';
	gitText.innerText = await installPC();
	installed = true;
	show();
});

uninstall.addEventListener('click', async () => {
	gitText.innerText = 'Uninstalling powercord...';
	gitText.innerText = await uninstallPC();
	installed = false;
	show();
});

plug.addEventListener('click', async () => {
	gitText.innerText = 'Plugging powercord...';
	gitText.innerText = await plugPowercord();
});

unplug.addEventListener('click', async () => {
	gitText.innerText = 'Unplugging powercord...';
	gitText.innerText = await unplugPowercord();
});

btnTheme.addEventListener('click', async () => {
	gitText.innerText = await downloadThemePlugin();
});

btnPlugin.addEventListener('click', async () => {
	gitText.innerText = await downloadPluginDownloader();
});

function show() {
	if (installed) divPlugins.classList.remove('disable');
	else divPlugins.classList.add('disable');
}

async function setup() {
	const cando = JSON.parse(await canDo());
	installed = await isInstalled();
	if (!cando.ok) {
		btnGit.disabled = true;
		uninstall.disabled = true;
		plug.disabled = true;
		unplug.disabled = true;
		gitText.innerHTML = `<a target=_blank href=${cando.link}>${cando.text}<a/>`;
	}
	show();
}

window.onload = setup;
