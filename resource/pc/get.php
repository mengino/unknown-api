<?php

class Spider
{
    public $arr = [];

    public function __construct()
    {

    }

    public function getResource(string $content, string $type): void
    {
        $pattern = '/(?<=' . ($type === 'js' ? 'src' : 'href') . '=")(.*?)\.' . $type . '?(\?(.*?))?(?=")/';
        preg_match_all($pattern, $content, $matches);

        foreach ($matches[0] as $url) {
            if (strpos($url, 'http') === false) {
                $url = 'https://www.homevv.com/static' . $url;
            }

            if (in_array($url, $this->arr)) {
                continue;
            }

            $file = end(explode('/', $url));
            $file = reset(explode('?', $file));

            !is_dir($type) && @mkdir($type);

            echo "download {$file} from {$url} \n";
            shell_exec("curl -s -o {$type}/{$file} {$url}");

            if (is_file("{$type}/{$file}") && filesize("{$type}/{$file}")) {
                $this->arr[] = $url;
            }
        }
    }

    public function inteadResourcePath(string $content): string
    {
        $content = str_replace('https://www.homevv.com', '', $content);
        $content = str_replace('/static', '', $content);

        return $content;
    }
}

$spider = new Spider();

$dir = scandir('./');

foreach ($dir as $value) {
    if (strpos($value, '.html') === false) {
        continue;
    }

    $name = reset(explode('.', $value));

    $content = file_get_contents("{$name}.html");

    $spider->getResource($content, 'js');
    $spider->getResource($content, 'css');

    $content = $spider->inteadResourcePath($content);

    file_put_contents("{$name}.html", $content);
}
