const btnInstall = document.querySelector('.install');
const text = document.querySelector('.text');
const uninstall = document.querySelector('.uninstall');
const btnPlug = document.querySelector('.porkPlugger');
const btnUnplug = document.querySelector('.porkUnplugger');
const btnTheme = document.querySelector('.themeDownloader');
const btnPlugin = document.querySelector('.pluginDownloader');
const btnUpdate = document.querySelector('.update');
const divPlugins = document.querySelector('.plugins');
let installed = false;

btnInstall.addEventListener('click', async () => {
	text.innerText = 'Installing powercord...';
	text.innerText = await installPC();
	installed = true;
	show();
});

uninstall.addEventListener('click', async () => {
	text.innerText = 'Uninstalling powercord...';
	text.innerText = await uninstallPC();
	installed = false;
	show();
});

btnPlug.addEventListener('click', async () => {
	text.innerText = 'Plugging powercord...';
	text.innerText = await plugPowercord();
});

btnUnplug.addEventListener('click', async () => {
	text.innerText = 'Unplugging powercord...';
	text.innerText = await unplugPowercord();
});

btnTheme.addEventListener('click', async () => {
	text.innerText = await downloadThemePlugin();
});

btnPlugin.addEventListener('click', async () => {
	text.innerText = await downloadPluginDownloader();
});

btnUpdate.addEventListener('click', async () => {
	text.innerText = await updatePowercord();
});

function show() {
	if (installed) {
		divPlugins.classList.remove('disable');
		uninstall.classList.remove('disable');
		btnInstall.classList.add('disable');
		btnPlug.classList.remove('disable');
		btnUnplug.classList.remove('disable');
		btnUpdate.classList.remove('disable');
	} else {
		divPlugins.classList.add('disable');
		uninstall.classList.add('disable');
		btnPlug.classList.add('disable');
		btnUnplug.classList.add('disable');
		btnInstall.classList.remove('disable');
		btnUpdate.classList.add('disable');
	}
}

async function setup() {
	const cando = JSON.parse(await canDo());
	installed = await isInstalled();
	if (!cando.ok) {
		btnInstall.disabled = true;
		uninstall.disabled = true;
		btnPlug.disabled = true;
		btnUnplug.disabled = true;
		text.innerHTML = `<a target=_blank href=${cando.link}>${cando.text}<a/>`;
	}
	show();
}

window.onload = setup;
