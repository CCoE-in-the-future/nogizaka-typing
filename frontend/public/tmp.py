from PIL import Image

def get_aspect_ratio(image_path):
    # 画像を開く
    with Image.open(image_path) as img:
        width, height = img.size
        return width, height

def crop_to_aspect_ratio(image_path, target_aspect_ratio):
    with Image.open(image_path) as img:
        width, height = img.size
        img_aspect_ratio = width / height
        
        if img_aspect_ratio > target_aspect_ratio:
            # 幅が大きすぎる場合
            new_width = int(height * target_aspect_ratio)
            left = (width - new_width) / 2
            right = left + new_width
            top = 0
            bottom = height
        else:
            # 高さが大きすぎる場合
            new_height = int(width / target_aspect_ratio)
            top = (height - new_height) / 2
            bottom = top + new_height
            left = 0
            right = width

        # トリミングを適用
        cropped_img = img.crop((left, top, right, bottom))
        return cropped_img

def process_images(image_paths):
    # 1枚目のアスペクト比を取得
    base_image_path = image_paths[0]
    base_aspect_ratio = get_aspect_ratio(base_image_path)
    target_aspect_ratio = base_aspect_ratio[0] / base_aspect_ratio[1]

    cropped_images = []
    for image_path in image_paths:
        # 各画像をトリミングして、ターゲットアスペクト比に合わせる
        cropped_img = crop_to_aspect_ratio(image_path, target_aspect_ratio)
        cropped_images.append([image_path, cropped_img])
    
    return cropped_images

# 画像ファイルのパスリスト
from glob import glob
image_paths = glob("*.jpeg")
print(f"{image_paths=}")
# image_paths = ["image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg", "image5.jpg"]

cropped_images = process_images(image_paths)

# トリミングした画像を保存する例
for i, img in enumerate(cropped_images):
    img[1].save(img[0])
