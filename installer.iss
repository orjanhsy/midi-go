[Setup]
AppName=MyApp
AppVersion=1.0
DefaultDirName={pf}\MyApp
OutputDir=.
DefaultGroupName=MyApp
OutputBaseFilename=myapp-installer
Compression=lzma
SolidCompression=yes

[Files]
Source: "myapp.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "libgcc_s_seh-1.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "libstdc++-6.dll"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{commondesktop}\MyApp"; Filename: "{app}\myapp.exe"
