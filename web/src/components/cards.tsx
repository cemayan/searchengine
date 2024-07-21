import {Card, CardHeader, CardFooter, Divider, Link, Image} from "@nextui-org/react";

interface CardsProps {
    title?: string
    url?: string
}


export const Cards = (props: CardsProps) => {
    return (
        <Card className="max-w-[400px]">
            <CardHeader className="flex gap-3">
                <Image
                    alt="nextui logo"
                    height={40}
                    radius="sm"
                    src="https://avatars.githubusercontent.com/u/86160567?s=200&v=4"
                    width={40}
                />
                <div className="flex flex-col">
                    <p className="text-md">{props.title}</p>

                </div>
            </CardHeader>
            <Divider/>
            <CardFooter>
                <Link
                    isExternal
                    showAnchorIcon
                    href={props.url}
                >
                    Visit
                </Link>
            </CardFooter>
        </Card>
    );
}
