use clap::{Args, Parser, Subcommand};
use std::{env, fs::File, io::Write};

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
struct Cli {
    #[structopt(subcommand)]
    command: Option<Commands>,
}

#[derive(Subcommand)]
enum Commands {
    #[clap(name = "util")]
    Util(Util),
    #[clap(name = "file")]
    File(AddFile),
}

#[derive(Args)]
struct AddFile {
    add: Option<String>,
}

#[derive(Parser)]
struct Util {
    #[structopt(subcommand)]
    command: UtilCommands,
}

#[derive(Subcommand)]
enum UtilCommands {
    #[clap(name = "reverse")]
    Reverse(Reverse),
    #[clap(name = "inspect")]
    Inspect(Inspect),
}

#[derive(Args)]
struct Reverse {
    string: Option<String>,
}

#[derive(Args)]
struct Inspect {
    string: Option<String>,
}

fn reverse(input: &String) -> String {
    input.chars().rev().collect()
}

fn inspect(input: &String) -> i32 {
    input.len() as i32
}

fn main() {
    let cli = Cli::parse();

    match &cli.command {
        Some(Commands::Util(util)) => match &util.command {
            UtilCommands::Reverse(string) => {
                let input = string.string.as_ref().unwrap();
                println!("{}", reverse(input));
            }
            UtilCommands::Inspect(string) => {
                let input = string.string.as_ref().unwrap();
                println!("{}", inspect(input));
            }
        },
        Some(Commands::File(file)) => match &file.add {
            Some(string) => {
                let current_dir = env::current_dir().unwrap();
                let mut file = match File::create(current_dir.join(string)) {
                    Ok(it) => it,
                    Err(err) => Err(err).unwrap(),
                };

                match file.write_all(b"Hello, there!") {
                    Ok(_) => println!("File created successfully"),
                    Err(err) => Err(err).unwrap(),
                }
            }
            None => println!("No add command provided"),
        },
        None => println!("No command provided"),
    }
}
