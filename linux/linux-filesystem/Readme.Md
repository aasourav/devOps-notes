It is possible to define two independent distinctions among files: 
- shareable vs. unshareable and 
- variable vs. static.

"Shareable" files are those that can be stored on one host and used on others. "Unshareable" files are those
that are not shareable. For example, the files in user home directories are shareable whereas device lock
files are not.

"Static" files include binaries, libraries, documentation files and other files that do not change without
system administrator intervention. "Variable" files are files that are not static


## /bin 

Contains essential command-line utilities required for system booting, maintenance, and repair


| Feature       | /bin                      | /usr/bin                          |
|---------------|---------------------------|-----------------------------------|
| **Purpose**   | Essential system utilities | Non-essential applications        |
| **Availability** | Always available           | Available after full system boot |
| **Mounting**  | Located in root `/` partition | May reside on a separate partition |
| **User Scope** | Root & system processes    | Regular users & developers        |
| **Examples**  | `cp`, `ls`, `cat`, `bash`   | `firefox`, `gcc`, `chsh`           |

### Why the Separation Exists
The primary reason for separating these directories is to support system resilience, maintenance, and flexibility.

### Who Uses /bin and /usr/bin
/bin commands are mainly used by system administrators and processes during booting and system recovery. Examples include:

- sh – essential shell for emergency use.
- mount – to mount filesystems.
- fsck – to check and repair file systems.

/usr/bin commands are used by regular users for everyday tasks and include:

- python – scripting and programming.
- gcc – compiler tools for development.
- xrandr – graphical display management tools.

some notes:
```md
POSIX (Portable Operating System Interface). POSIX ensures compatibility and portability across different Unix-like operating systems by specifying how programs should interact with the operating system, including aspects such as file operations, process management, and shell scripting.

Why is POSIX Compliance Important?

Portability:
Programs and scripts written on one POSIX-compliant system (e.g., Linux) will work on others (e.g., macOS, FreeBSD) without modification.

Interoperability:
Provides a common foundation for software developers, ensuring software can run across various Unix-like systems.

Reliability and Consistency:
POSIX-compliant systems offer predictable behavior across different environments, which is crucial for enterprise applications.

Cross-Platform Development:
Developers can target multiple operating systems without platform-specific changes, reducing development effort.

```


## /boot

Holds essential files for booting (kernel, bootloader files, initial RAM disk).

/sbin - ontains executable programs for managing the boot process.
/etc - Contains bootloader configuration files that aren't needed during the boot process.



[Book Download Link](https://d1wqtxts1xzle7.cloudfront.net/101343656/fhs-3.0-libre.pdf?1682104856=&response-content-disposition=attachment%3B+filename%3DFilesystem_hierarchy_standard.pdf&Expires=1737521442&Signature=alx49QKznyEoGlb2fQh4OgXHtlTpeos6m644Ki-a07KT~i~JcBaZo~mF~woxaVQEqYZnXmKtCcBrhBQziyPB0kufa7BM2yn8Ibdc8GnCnnH2WM~B0XgIdbQPHqcP3qqbt3~JIkRWbX8AN-vCgkvklZRr3GeEMeSx~tUegmn4eo8a2y0HuOkMceXzXyWOyIWPf736qmo7WuYAcXhSlxz4ugCpM-6XE66YH7GITiFZBXaBGoy0II1np6S~F3Ev6UHiUHEkBtGLK4GgZ4vUPKEndjsKJ7ZcL02pmZ-oyyb3jNgAdufMDgiyVH5YlupIuA97ELLY4e3cyOJ7MGdvu0mauw__&Key-Pair-Id=APKAJLOHF5GGSLRBV4ZA)