Co má program dělat?
Jednoduše vidět stoly po zavolání, jaká data jsou volná, jak jsou velké - done
musí umožnit spojit stoly
následně napojit na TG
Jednoduché UI pro bookování je cíl, tj ideálně klikačka (webové)
běžně by měl ukázat 7 dní dopředu
dále by mohl sbírat data, jak moc byl podnik využit

Momentálně používají tg skupinu na rezervace
Mělo by možné být volat i textově, co chci

Mimo to bude dělat i webové UI pro obě možnosti

Co to nemá dělat - obtižnější práci - tj není nutné psát přesný počet lidí
Může to dělat grafy jak které hodiny je podnik obsazen


Muze to vygenerovat mapku stolu - zelene a cervene dle obsazenosti

Prvni fáze
Udelat tui - kde napisu rezervace pocet lidi a navrhne mi to mozne volby - cca semi done?
Nasledne rozchodit hodiny - semi done

Melo by to mit osetrenou otviraci dobu
Momentálně se přesouvám k mongodb místo jsonu

Mělo by to umět udělat first setup - to je teď je vše o manuální přidání věcí, to bude asi udělané v admin prostředí co se týče jmén a kapacit, ale bude potřeba určitě udělat setup, který udělá inicialní nastavení - to je běžím tady, stáhnout docker a nastavit - možná to celé dát do docker compose? 

Nebylo by špatné mít admin rozhraní, kde by se snadno spravovala databáze stolů, popřípadě nadřazeně měnili rezervace a tak

Prvně tedy udělám nějaké stránky pomoci hugo - snadno a rychle a hlavně rychlé

mongodb funguje skrze docker

NICE to have: ruzné další napojení na jiné služby
dockerizace
exporty
